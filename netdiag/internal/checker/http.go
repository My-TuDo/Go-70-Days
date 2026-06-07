package checker

import (
	"fmt"
	"net/http"
	"time"
)

// CheckHTTP 对 url 做 HTTP GET 请求
// 成功 → Status="✅", Detail="HTTP 200"
// 失败 → Status="❌", Error=错误信息
//
// 提示：
//   - 用 http.Client{Timeout: 5 * time.Second}
//   - resp.StatusCode 获取状态码
//   - 记得 defer resp.Body.Close()
func CheckHTTP(url string) Result {
	// TODO: 实现 HTTP 探测

	// 记录开始时间
	start := time.Now()

	// 创建 HTTP 客户端并设置超时时间
	client := http.Client{Timeout: 5 * time.Second}

	// 发送 GET 请求
	resp, err := client.Get(url)

	// 计算请求耗时
	latency := time.Since(start).Milliseconds()

	// 请求失败
	if err != nil {
		return Result{
			Name:      "HTTP 检查",
			Target:    url,
			Status:    "❌",
			Error:     err.Error(),
			LatencyMs: latency,
		}
	}
	defer resp.Body.Close()

	// 请求成功
	return Result{
		Name:      "HTTP 检查",
		Target:    url,
		Status:    "✅",
		Detail:    fmt.Sprintf("HTTP %d", resp.StatusCode), // 其他状态码也算成功，但要显示状态码
		LatencyMs: latency,
	}
}
