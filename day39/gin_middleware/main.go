package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

/*
*	Gin 中间件:像是一个滤网，可以在请求到达路由处理函数之前或之后执行一些操作
*	可以利用中间件来实现统一日志记录、异常恢复、QPS限流
 */

// StatCost 是一个自定义的监控中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1.请求进来时
		start := time.Now()

		// 这里可以设置一个请求指纹，方便全链路追踪
		c.Set("trace_id", "trace-123456")

		// 2.执行后续逻辑
		// c.Next() 的意思是：放行，让后面的业务逻辑跑起来
		c.Next()

		// 3.业务逻辑跑完，响应出去时
		// 计算耗时
		cost := time.Since(start)
		path := c.Request.URL.Path
		status := c.Writer.Status()

		fmt.Printf("[监控报表]路径：%s | 状态：%d | 耗时：%v\n", path, status, cost)

		// 如果耗时超过500ms， 在真实场景中我们可能会报警，或者记录到监控系统中
		if cost > 500*time.Millisecond {
			fmt.Println("警告：检测到慢接口响应！")
		}
	}
}

func main() {
	r := gin.New() // 使用New而不是Default， 这样我们就不会默认使用Logger和Recovery中间件了，可以自己定制中间件

	// 注册哨兵中间件
	// 这样以后所有的路由都会被这个逻辑包裹
	r.Use(StatCost())
	r.Use(gin.Recovery()) // 崩溃自动回复中间件，防止服务器宕机

	r.GET("/api/v1/ping", func(c *gin.Context) {
		// 模拟业务逻辑
		trace, _ := c.Get("trace_id")
		c.JSON(200, gin.H{
			"message": "pong",
			"trace":   trace,
		})
	})

	r.GET("/api/v1/slow-job", func(c *gin.Context) {
		// 模拟一个慢任务
		time.Sleep(600 * time.Millisecond)
		c.JSON(200, gin.H{"status": "done"})
	})

	fmt.Println("性能监控网关已启动在：8080")
	r.Run(":8080")

}
