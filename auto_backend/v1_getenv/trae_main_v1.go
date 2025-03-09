package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

func main() {
	// 初始化数据库
	var err error
	db, err = gorm.Open(sqlite.Open("./todo.db"), &gorm.Config{})
	if err != nil {
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
