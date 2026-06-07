package checker

import (
	"crypto/tls"
	"fmt"
	"time"
)

// CheckTLS 对 host 做 TLS 握手，检查证书
// 成功 → Status="✅", Detail=证书过期时间
// 失败 → Status="❌", Error=错误信息
//
// 提示：
//   - 用 tls.DialWithDialer 或 &tls.Config{InsecureSkipVerify: true}
//   - conn.ConnectionState().PeerCertificates[0].NotAfter 获取证书过期时间
//   - 记得 conn.Close()
func CheckTLS(host string) Result {
	// TODO: 实现 TLS 证书检查

	start := time.Now()

	// 创建 TLS 连接
	conn, err := tls.Dial("tcp", host+":443", &tls.Config{InsecureSkipVerify: true})

	latency := time.Since(start).Milliseconds()

	if err != nil {
		return Result{
			Name:      "TLS 证书",
			Target:    host,
			Status:    "❌",
			Error:     err.Error(),
			LatencyMs: latency,
		}
	}
	defer conn.Close()

	// 获取证书过期时间
	cert := conn.ConnectionState().PeerCertificates[0].NotAfter
	return Result{
		Name:      "TLS 证书",
		Target:    host,
		Status:    "✅",
		Detail:    fmt.Sprintf("证书过期时间: %s", cert.Format("2006-01-02 15:04:05")),
		LatencyMs: latency,
	}
}

/*
* conn.ConnectionState().PeerCertificates[0].NotAfter 获取证书过期时间
* 从右往左读
*	NotAfter <- 这个证书是什么时候过期的？
*	PeerCertificates[0] <- 第一个证书是哪个？
*	.ConnectionState() <- 这个连接的当前状态是什么？
*	conn <- TLS连接对象
 */
