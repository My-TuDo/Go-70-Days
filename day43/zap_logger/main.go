package main

import (
	"time"

	"go.uber.org/zap" // [核心包]： Uber 出品的高性能日志库
)

func main() {
	// [工程环境]： 模拟一个服务器的生产启动配置

	// 1. [zap.NewPriduction]
	// 作用：创建一个预设好的生产环境日志器
	// 特点：默认输出 JSON 格式、包含时间戳、日志级别、调用行号
	// 意义：这种格式可以直接被 Logstash 采集
	logger, _ := zap.NewProduction()

	// 2. [defer logger.Sync()]
	// 作用：刷新缓冲区
	// 职业修养：日志通常三先写在内存缓冲区的，程序退出前必须刷入硬盘，否则会丢失最后几条日志
	defer logger.Sync()

	// 3. [logger.With(...)]
	// 作用：创建一个带“基础上下文”的日志器
	// 场景：比如这一批任务都属于“Job-999”，这样后面的每行日志都会加上这个ID
	jobLogger := logger.With(
		zap.String("service", "sign-worker"),
		zap.Int("cluster_id", 7),
	)

	jobLogger.Info("自动化运维任务开始执行") // 输出一条 INFO 级别的日志

	// 模拟一次网络探测
	target := "api-server.com"
	start := time.Now()

	// 假设这里发生了一个错误
	isSuccess := false

	if !isSuccess {
		// 4. [logger.Error(...)]
		// 作用：记录错误级别日志
		// 细节：不要用字符串拼接，要用 zap.String 这种强类型字段
		// 优势：避免了接口反射开销，速度极快
		jobLogger.Error("网络探测失败",
			zap.String("target", target),
			zap.Duration("latency", time.Since(start)),
			zap.String("reason", "connection_reset"),
		)
	} else {
		jobLogger.Info("网络探测成功",
			zap.Duration("cost", time.Since(start)),
		)
	}
}
