package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 模型定义
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" binding:"required" gorm:"uniqueIndex"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Job      string `json:"job"`
	Level    string `json:"level"` // 编程语言熟练程度
	Lang     string `json:"lang"`  // 擅长的开发语言
}

// 配置结构体，这个结构很离谱，所以还是得用viper
type Config struct {
	Database struct {
		Type  string `json:"type"`
		MySQL struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Host     string `json:"host"`
			Port     string `json:"port"`
			DBName   string `json:"dbname"`
		} `json:"mysql"`
		SQLite struct {
			Path string `json:"path"`
		} `json:"sqlite"`
	} `json:"database"`
}

var db *gorm.DB
var config Config

// 加载配置文件
func loadConfig() error {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println("无法读取配置文件:", err)
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Println("解析配置文件失败:", err)
		return err
	}

	return nil
}

// 初始化数据库连接
func initDatabase() error {
	var err error

	// 尝试加载配置文件
	configErr := loadConfig()

	// 确定数据库类型
	dbType := ""
	if configErr == nil {
		dbType = strings.ToLower(config.Database.Type)
	}

	// 如果配置文件中没有指定或加载失败，则从环境变量获取
	if dbType == "" {
		dbType = strings.ToLower(os.Getenv("DB_TYPE"))
	}

	switch dbType {
	case "mysql":
		// MySQL连接配置
		username := ""
		password := ""
		host := "localhost"
		port := "3306"
		dbname := "trae_demo"

		// 优先使用配置文件中的设置
		if configErr == nil {
			if config.Database.MySQL.Username != "" {
				username = config.Database.MySQL.Username
			}
			if config.Database.MySQL.Password != "" {
				password = config.Database.MySQL.Password
			}
			if config.Database.MySQL.Host != "" {
				host = config.Database.MySQL.Host
			}
			if config.Database.MySQL.Port != "" {
				port = config.Database.MySQL.Port
			}
			if config.Database.MySQL.DBName != "" {
				dbname = config.Database.MySQL.DBName
			}
		}

		// 如果配置文件中没有设置或加载失败，则从环境变量获取
		if username == "" {
			username = os.Getenv("DB_USERNAME")
		}
		if password == "" {
			password = os.Getenv("DB_PASSWORD")
		}
		envHost := os.Getenv("DB_HOST")
		if envHost != "" {
			host = envHost
		}
		envPort := os.Getenv("DB_PORT")
		if envPort != "" {
			port = envPort
		}
		envDBName := os.Getenv("DB_NAME")
		if envDBName != "" {
			dbname = envDBName
		}

		dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		// 默认使用SQLite
		dbPath := ""

		// 优先使用配置文件中的设置
		if configErr == nil && config.Database.SQLite.Path != "" {
			dbPath = config.Database.SQLite.Path
		} else {
			// 如果配置文件中没有设置或加载失败，则从环境变量获取
			dbPath = os.Getenv("DB_PATH")
			if dbPath == "" {
				dbPath = "./todo.db"
			}
		}

		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	}

	return err
}

func main() {
	// 初始化数据库
	if err := initDatabase(); err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// 自动迁移数据库结构
	db.AutoMigrate(&User{})

	// 创建Gin路由
	r := gin.Default()

	// 配置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// API路由组
	api := r.Group("/api")
	{
		// 用户相关路由
		users := api.Group("/users")
		{
			users.GET("", getUsers)
			users.GET("/:id", getUserByID)
			users.POST("", createUser)
			users.PUT("/:id", updateUser)
			users.DELETE("/:id", deleteUser)
		}
	}

	// 启动服务器
	log.Println("Server starting on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// 获取所有用户
func getUsers(c *gin.Context) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户列表失败"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// 根据ID获取用户
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User

	result := db.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// 创建新用户
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// 更新用户
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	// 检查用户是否存在
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	// 绑定请求体到用户结构
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新用户
	db.Save(&user)

	c.JSON(http.StatusOK, user)
}

// 删除用户
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	// 检查用户是否存在
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	// 删除用户
	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "用户已删除"})
}
