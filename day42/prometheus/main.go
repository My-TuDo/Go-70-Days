package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"          // 定义监控指标的核心包
	"github.com/prometheus/client_golang/prometheus/promhttp" // 暴露监控数据的 HTTP 处理包
)

var (
	// 1. [prometheus.NewCounter]: 定义一个自增计数器
	// 工程意义：用于记录 “累积发生的事件”， 通常配合 Grafana 的 rate() 函数看每秒 QPS
	requestCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "api_requests_total",
		Help: "模拟接口请求总数",
	})

	// 2. [prometheus.NewGauge]: 定义一个 Gauge 类型的指标
	// 作用： 定义一个压力表， 数值可以上下浮动
	// 工程意义： 用于监控 “瞬时状态”， 比如服务器当前的 CPU 温度或者在线人数
	activePlayers = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "active_players_count",
		Help: "模拟在线玩家数量",
	})
)

func init() {
	// 3. [prometheus.MustRegister]: 将定义好的指标注册到 Prometheus 的默认注册器中
	// 作用： 让 Prometheus 知道我们定义了哪些指标， 以便它们能够被正确地收集和暴露
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(activePlayers)
}

func main() {
	// [工程环境]： 模拟一个繁忙的游戏后段网关

	// 模拟业务逻辑：每秒随机产生一些流量
	go func() {
		for {
			requestCount.Inc()                                           // 每次请求计数器加1
			activePlayers.Set(rand.Float64() * 1000)                     // 在线玩家数随机在0-1000之间浮动
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // 模拟请求间隔
		}
	}()

	// 4. [promhttp.Handler]
	// 作用： 这是一个特殊的网页接口。
	// 工程意义： Prometheus 服务器会每隔 15 秒访问这个接口，把数据拉出来
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("监控指标已暴露在 http://localhost:8081/metrics")

	http.ListenAndServe(":8081", nil) // 启动 HTTP 服务器，监听 8081 端口
}
