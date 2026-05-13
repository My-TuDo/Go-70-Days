package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

// 1.[全局变量声明]：监控指标
// 工程意义：这些变量将直接决定 Grafana 看板上看到的曲线
var (
	probeTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sre_probe_total",
			Help: "累积探测次数",
		},
		[]string{"target", "status"}, // 标签思维：按目标和状态分类统计
	)
)

func init() {
	prometheus.MustRegister(probeTotal)
}

// 2.[核心业务逻辑]：拨测引擎
// 作用：探测目标地址是否存活，并集成日志与指标
func runProbe(ctx context.Context, logger *zap.Logger, target string) {
	// 3.[net.Dialer]:底层拨号器
	dialer := &net.Dialer{}

	// 4.[DialContext]:结合 Context 的探测
	// 意义：如果网络卡死，Context 会在规定时间强制切断 Socket，防止产生“僵尸连接”
	start := time.Now()
	conn, err := dialer.DialContext(ctx, "tcp", target)

	duration := time.Since(start)

	if err != nil {
		// 5.[日志记录]：失败路径
		logger.Error("探测失败",
			zap.String("target", target),
			zap.Error(err),
			zap.Duration("latency", duration),
		)
		probeTotal.WithLabelValues(target, "failed").Inc()
		return
	}
	defer conn.Close()

	// 6.[日志记录]：成功路径
	logger.Info("探测成功",
		zap.String("target", target),
		zap.Duration("latency", duration),
	)
	probeTotal.WithLabelValues(target, "success").Inc()
}

func main() {
	// 7.[日志器初始化]
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// 8.[启动监控网关]（Prometheus抓取入口）
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8081", nil)
	}()

	fmt.Println("已启动...")

	// 9.[主巡检循环]
	targets := []string{"api-server.com:443", "8.8.8.8:53", "127.0.0.1:22"}

	for {
		for _, target := range targets {
			// 10.[超时控制]：给每一次探测 3 秒的时间限制
			// 保护程序自身不挂掉
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

			runProbe(ctx, logger, target)

			cancel() // 释放资源
		}
		time.Sleep(10 * time.Second) // 每 10 秒巡检一次
	}
}
