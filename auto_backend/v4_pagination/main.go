package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

var db *gorm.DB

// 加载配置文件
func loadConfig() error {
	// 设置配置文件名
	viper.SetConfigName("config")
	// 设置配置文件类型
	viper.SetConfigType("json")
	// 设置配置文件路径
	viper.AddConfigPath(".")

	// 设置环境变量前缀，用于从环境变量读取配置
	viper.SetEnvPrefix("DB")
	// 将.替换为_，使嵌套配置可以通过环境变量设置
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// 绑定环境变量
	viper.AutomaticEnv()

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("无法读取配置文件:", err)
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
	dbType := viper.GetString("database.type")

	// 如果配置文件中没有指定或加载失败，则使用默认值
	if dbType == "" || configErr != nil {
		dbType = "sqlite" // 默认使用SQLite
	}

	dbType = strings.ToLower(dbType)

	// debug print
	mongodb_port := viper.GetInt("database.mongodb.port")
	fmt.Println("[D] mongodb port: ", mongodb_port)

	switch dbType {
	case "mysql":
		// MySQL连接配置
		username := viper.GetString("database.mysql.username")
		password := viper.GetString("database.mysql.password")
		host := viper.GetString("database.mysql.host")
		if host == "" {
			host = "localhost"
		}
		port := viper.GetString("database.mysql.port")

		if port == "" {
			port = "3306"
		}
		dbname := viper.GetString("database.mysql.dbname")
		if dbname == "" {
			dbname = "trae_demo"
		}

		dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		// 默认使用SQLite
		dbPath := viper.GetString("database.sqlite.path")
		if dbPath == "" {
			dbPath = "todo.db" // 默认数据库文件路径
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
	// 获取分页参数
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	// 获取每页显示数量，限制在允许的范围内：5、10、20、100
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		pageSize = 10
	} else {
		// 确保pageSize是允许的值之一
		allowedSizes := map[int]bool{5: true, 10: true, 20: true, 50: true, 100: true}
		if !allowedSizes[pageSize] {
			pageSize = 10 // 默认值
		}
	}

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询总记录数
	var total int64
	db.Model(&User{}).Count(&total)

	// 查询当前页数据
	var users []User
	result := db.Offset(offset).Limit(pageSize).Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户列表失败"})
		return
	}

	// 返回数据，包含分页信息
	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"pagination": gin.H{
			"total":    total,
			"page":     page,
			"pageSize": pageSize,
			"pages":    int(math.Ceil(float64(total) / float64(pageSize))),
		},
	})
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
