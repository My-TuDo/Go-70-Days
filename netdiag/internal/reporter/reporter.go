package reporter

import (
	"encoding/json"
	"fmt"
	"strings"

	"netdiag/internal/checker"
)

// PrintSummary 在终端打印诊断结果
//
// 格式示例：
//
//	=== 网络诊断报告 ===
//	目标: example.com
//	────────────────────────
//	✅ DNS 解析   | example.com | 93.184.216.34 (12ms)
//	✅ TCP 检查   | example.com:443 | 端口开放 (45ms)
//	✅ HTTP 检查  | https://example.com | HTTP 200 (120ms)
//	✅ TLS 证书   | example.com | 2026-08-01 过期 (30ms)
func PrintResults(results []checker.Result) {
	// TODO: 打印格式化的诊断报告
	// 用 fmt.Println 和 fmt.Printf
	fmt.Println("=== 网络诊断报告 ===")
	for _, r := range results {
		fmt.Println(r.String())
	}
	fmt.Println(strings.Repeat("─", 50))
}

// SaveJSON 将结果保存为 JSON 文件（输出到终端，保持简单）
//
// 提示：json.MarshalIndent(results, "", "  ")
func SaveJSON(results []checker.Result) (string, error) {
	// TODO: 序列化成 JSON 字符串返回

	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
