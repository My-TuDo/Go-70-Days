package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1. 定义监控报告的结构
type HealthReport struct {
	Status    string  `json:"status"`    // 运行状态
	CPUUsage  float64 `json:"cpu_usage"` // CPU使用率
	MemUsage  float64 `json:"mem_usage"` // 内存使用率
	DBStatus  string  `json:"db_status"` // 数据库连接状态
	Timestamp string  `json:"timestamp"` // 时间戳
}

var DB *gorm.DB

// initDatabase 数据库初始化 (Database Initialization)
func initDatabase() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_70_days?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败（DB Connection Failed）")
	}
}

// getSystemStats 获取系统指标 (Get System Statistics)
func getSystemStats() (float64, float64) {
	// 获取CPU使用率
	cpuPercent, _ := cpu.Percent(time.Second, false)
	// 获取内存使用率
	memStats, _ := mem.VirtualMemory()
	// 返回第一个核心的占用和总占用比
	return cpuPercent[0], memStats.UsedPercent
}

func main081() {
	initDatabase()

	DB.AutoMigrate(&HealthReport{})

	r := gin.Default()

	// 运维监控接口 (Monitoring API)
	r.GET("/api/vl/health", func(c *gin.Context) {
		cpuUsage, memUsgae := getSystemStats()

		// 检查数据库是否活着
		dbStatus := "Running"
		sqlDB, err := DB.DB()
		if err != nil || sqlDB.Ping() != nil {
			dbStatus = "Down"
		}

		// 组装报告
		report := HealthReport{
			Status:    "OK",
			CPUUsage:  cpuUsage,
			MemUsage:  memUsgae,
			DBStatus:  dbStatus,
			Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		}
		DB.Create(&report)
		c.JSON(http.StatusOK, report)
	})

	fmt.Println(" Gopher 运维哨兵已启动...")
	r.Run(":8080")
}
