package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1. 定义模型(Model)
type Users struct {
	ID   uint   `gorm:"primaryKey" json:"id"` // 主键
	Name string `json:"name"`                 // 用户名
	Age  int    `json:"age"`                  // 年龄
}

// 2. 定义全局数据库对象
var DB *gorm.DB

// 3. 初始化数据库连接
func initDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_70_days?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败！" + err.Error())
	}
	fmt.Println("成功连接到数据库！")
	DB.AutoMigrate(&Users{})
	fmt.Println("数据库表结构同步完成！")
}

func main073() {
	// 第一步： 启动时连接数据库
	initDB()

	// 第二步： 初始化Web引擎
	r := gin.Default()

	// 第三步： 定义接口
	// 接口A： 增加用户（尝试访问： http://localhost:8080/add?name=张三&age=20）
	r.GET("/add", func(c *gin.Context) {
		// 从网址参数获取数据
		name := c.Query("name")
		ageStr := c.DefaultQuery("age", "18") // 如果age参数没有提供，默认值为18

		// 2. [核心] 字符串转整型
		// Atoi 的意思是 “ASCII to Integer”
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			age = 18
		}

		// 创建结构体对象
		newUser := Users{Name: name, Age: age} // 这里可以做字符串转整型，暂简单处理

		// GORM插入数据
		result := DB.Create(&newUser)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "存储失败"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "用户已存入数据库！",
			"user":    newUser,
		})
	})

	// 接口B： 查询所有用户（尝试访问： http://localhost:8080/users）
	r.GET("/users", func(c *gin.Context) {
		var users []Users
		// GORM查询所有数据并存入切片
		DB.Find(&users)

		// 吐出JSON给浏览器
		c.JSON(http.StatusOK, users)
	})

	// 第四步： 启动服务器
	r.Run(":8080")
}
