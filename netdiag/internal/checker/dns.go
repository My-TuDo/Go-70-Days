package checker

import (
	"net"
	"time"
)

// CheckDNS 对 host 做 DNS 解析
// 成功 → Status="✅", Detail=第一个 IP
// 失败 → Status="❌", Error=错误信息
func CheckDNS(host string) Result {
	start := time.Now()
	ips, err := net.LookupHost(host)
	latency := time.Since(start).Milliseconds()

	result := Result{
		Name:      "DNS 解析",
		Target:    host,
		LatencyMs: latency,
	}

	if err != nil {
		result.Status = "❌"
		result.Error = err.Error()
		return result
	}

	result.Status = "✅"
	result.Detail = ips[0]
	return result
}
