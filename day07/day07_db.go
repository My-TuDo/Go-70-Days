package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1. 定义模型(Model)
// GORM会自动把这个结构体变成一张叫`users`的表
type User struct {
	ID        uint      `gorm:"primaryKey"` // 主键
	Name      string    `gorm:"size:255"`   // 用户名
	Age       int       // 年龄
	CreatedAt time.Time // 创建时间
}

func main() {
	// 2.数据库连接字符串（DSN）
	// 格式： 用户名：密码@tcp(地址：端口)/数据库名?参数
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_70_days?charset=utf8mb4&parseTime=True&loc=Local"

	// 3.连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败！" + err.Error())
	}
	fmt.Println("成功连接到数据库！")

	// 4.自动迁移（AutoMigrate）
	// 这行代码会自动在 MySQL 里创建 users 表， 不需要手写 SQL 建表语句
	db.AutoMigrate(&User{})
	fmt.Println("数据库表结构同步完成！")

	// 5. 插入一条数据（增）
	newUser := User{Name: "Gopher大佬", Age: 21}
	db.Create(&newUser)
	fmt.Println("成功插入一条玩家数据！")

	// 6. 查询数据（查）
	var result User
	db.First(&result, "name = ?", "Gopher大佬") // 根据名字查询
	fmt.Printf("查询结果：ID=%d, Name=%s, Age=%d\n", result.ID, result.Name, result.Age)
}
