package checker

import (
	"net"
	"time"
)

// CheckTCP 对 host:port 做 TCP 连接检查
// 成功 → Status="✅", Detail="端口开放"
// 失败 → Status="❌", Error=错误信息
//
// 提示：
//   - 用 net.DialTimeout("tcp", host+":"+port, 5*time.Second)
//   - 连接成功记得 conn.Close()
func CheckTCP(host, port string) Result {
	// TODO: 实现 TCP 端口检查

	// 记录开始时间
	start := time.Now()

	// 尝试连接 TCP
	conn, err := net.DialTimeout("tcp", host+":"+port, 5*time.Second)

	// 计算连接耗时
	latency := time.Since(start).Milliseconds()

	// 连接失败
	if err != nil {
		return Result{
			Name:      "TCP 检查",
			Target:    host + ":" + port,
			Status:    "❌",
			Error:     err.Error(),
			LatencyMs: latency,
		}
	}
	conn.Close()

	// 连接成功
	return Result{
		Name:      "TCP 检查",
		Target:    host + ":" + port,
		Status:    "✅",
		Detail:    "端口开放",
		LatencyMs: latency,
	}
}
