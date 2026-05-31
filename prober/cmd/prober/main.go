package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"prober/internal/collector"
	"prober/internal/handler"
	"prober/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 程序启动流程：
//
//  1. 读取 config.yaml 获取目标列表
//  2. 连接 MySQL（自动建表）
//  3. 连接 Redis
//  4. 启动 HTTP 服务（Gin）
//  5. 启动后台拨测循环（每 30 秒执行一次）
//  6. 监听退出信号，优雅关闭
//
// 你需要实现以下 TODO 标记的部分

func main() {
	// ===== 1. 读取配置 =====
	config := loadConfig("config.yaml")
	fmt.Printf("已加载 %d 个探测目标\n", len(config.Targets))

	// ===== 2. 连接 MySQL =====
	// TODO: 用 GORM 连接 MySQL
	// dsn := "root:123456@tcp(127.0.0.1:3306)/prober?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	//     log.Fatalf("连接 MySQL 失败: %v", err)
	// }
	// 自动迁移（建表）
	// db.AutoMigrate(&model.ProbeResult{})
	// fmt.Println("MySQL 连接成功")

	var db *gorm.DB // 先占位，等你填
	_ = db

	// ===== 3. 连接 Redis =====
	// TODO: 用 redis.NewClient 连接 Redis
	// rdb := redis.NewClient(&redis.Options{
	//     Addr: "localhost:6379",
	// })
	// fmt.Println("Redis 连接成功")

	var rdb *redis.Client // 先占位，等你填
	_ = rdb

	// ===== 4. 启动 HTTP 服务 =====
	r := gin.Default()
	handler.SetupRoutes(r, db, rdb, config.Targets)

	// 在后台启动 HTTP 服务
	go func() {
		fmt.Println("HTTP 服务已启动: http://localhost:8080")
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("HTTP 服务启动失败: %v", err)
		}
	}()

	// ===== 5. 启动拨测循环 =====
	// 先立即执行一次，然后每 30 秒执行一次
	go func() {
		// 立即执行第一轮
		collector.RunProbeCycle(config.Targets, db)

		// 每 30 秒执行一次
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			collector.RunProbeCycle(config.Targets, db)
		}
	}()

	// ===== 6. 优雅退出 =====
	// TODO: 监听 SIGINT/SIGTERM 信号
	// 收到信号后打印退出信息并结束程序
	// 参考 day20/graceful
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("\n正在退出...")
}

// loadConfig 读取 YAML 配置文件
func loadConfig(path string) model.ProbeConfig {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	var config model.ProbeConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	return config
}
