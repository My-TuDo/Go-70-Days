package handler

import (
	"fmt"
	"net/http"

	"prober/internal/collector"
	"prober/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 提示：你需要用到的知识点
//
// gin.Context — 处理 HTTP 请求和响应（参考 day07、day39）
// c.JSON — 返回 JSON 响应
// c.Query — 获取 URL 查询参数 /history?name=baidu

// SetupRoutes 注册所有路由
//
// 你需要注册三个路由：
//
//	GET /health     → 返回 {"status": "ok"}
//	GET /status     → 返回所有目标的最新状态（从 Redis 读）
//	GET /history    → 返回指定目标的历史记录（从 MySQL 读，需 ?name=xxx 参数）
func SetupRoutes(r *gin.Engine, db *gorm.DB, rdb *redis.Client, targets []model.TargetConfig) {
	// TODO: 注册三个路由
	// r.GET("/health", func(c *gin.Context) { ... })
	// r.GET("/status", ...)
	// r.GET("/history", ...)
	//
	// 提示：路由处理函数写在下面三个单独的函数里

	r.GET("/health", HealthHandler)
	r.GET("/status", func(c *gin.Context) {
		StatusHandler(c, rdb, targets) // main.go 中已经 loadConfig
	})
	r.GET("/history", func(c *gin.Context) {
		HistoryHandler(c, db)
	})
}

// HealthHandler 返回服务健康状态
//
// 你需要实现：返回 JSON {"status": "ok"}
func HealthHandler(c *gin.Context) {
	// TODO: c.JSON(http.StatusOK, gin.H{"status": "ok"})
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// StatusHandler 从 Redis 读取所有目标的最新状态并返回
//
// 因为 Redis 是按 key 存的，你需要遍历 targets 逐个读取
//
// 你需要实现：
//  1. 遍历所有 targets
//  2. 对每个 target 调用 collector.GetCachedStatus
//  3. 收集所有结果，返回 JSON 数组
//
// 提示：targets 可以从哪里获取？
//
//	你可以把 config 也传进来，或者重新从文件读取
//	这里简化处理：把 targets 作为全局变量或参数传进来
func StatusHandler(c *gin.Context, rdb *redis.Client, targets []model.TargetConfig) {
	// TODO: 这里需要 targets 列表才能遍历读取缓存
	// 目前先返回一个示意，等你实现了 model 和 config 读取后再完善

	// 创建收集器，收集获取到的 cache 数据
	var statusList []*model.ProbeCache

	for _, target := range targets {
		cache, err := collector.GetCachedStatus(rdb, target.Name)
		if err != nil {
			fmt.Printf("读取缓存失败: %v\n", err)
			continue
		}
		if cache != nil {
			// 处理获取到的缓存数据
			fmt.Printf("目标: %s, 状态: %v\n", target.Name, cache)
			statusList = append(statusList, cache)
		}
	}
	c.JSON(http.StatusOK, statusList)
}

// HistoryHandler 从 MySQL 查询某个目标的历史记录
//
// URL 参数: ?name=baidu
// 返回最近 20 条记录，按时间倒序
//
// 你需要实现：
//  1. 用 c.Query("name") 获取目标名称
//  2. 如果 name 为空，返回错误
//  3. 用 db.Where("target_name = ?", name).Order("created_at DESC").Limit(20).Find(&results)
//  4. 返回 JSON
func HistoryHandler(c *gin.Context, db *gorm.DB) {
	// TODO: 在这里写你的代码
	//

	// 从 URL 查询参数获取目标名称
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 name 参数"})
		return
	}

	var results []model.ProbeResult
	db.Where("target_name = ?", name).Order("created_at desc").Limit(20).Find(&results)
	c.JSON(http.StatusOK, results)

	// name := c.Query("name")
	// if name == "" {
	//     c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 name 参数"})
	//     return
	// }
	//
	// var results []model.ProbeResult
	// db.Where("target_name = ?", name).Order("created_at desc").Limit(20).Find(&results)
	//
	// c.JSON(http.StatusOK, results)

}
