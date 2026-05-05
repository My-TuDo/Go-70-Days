package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

/*
*	Graceful Exit:优雅退出，指的是在程序接收到终止信号时，能够有序地关闭资源、完成正在处理的请求，并最终安全地退出程序的过程
*	场景：CI/CD流水线：自动发布新版本时；系统维护：需要重启宿主机或迁移容器时；自动扩缩容：流量低谷期关闭多余的服务器节点时
*
 */

func main() {
	r := gin.Default()

	// 模拟一个长耗时任务（比如正在上传大文件）
	r.GET("long-task", func(c *gin.Context) {
		// 模拟处理需要5秒钟
		time.Sleep(5 * time.Second)
		c.JSON(200, gin.H{"message": "任务完成！"})
	})

	// 1. [核心改动]：不要直接 r.Run()，要手动构建 server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// 2. [核心改动]：开启一个协程去启动服务器
	// 这样主协程才能空出来监听退出信号
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("监听失败：%s\n", err)
		}
	}()

	// 3. [核心改动]：设置信号拦截器
	// 监听 Ctrl+C(SIGINT) 或者 kill(SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 程序i会在这里阻塞，直到收到退出信号
	<-quit
	log.Println("收到退出信号，正在优雅关闭服务器...")
}
