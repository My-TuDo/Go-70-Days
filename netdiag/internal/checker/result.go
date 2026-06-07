package checker

import "fmt"

// Result 每个检查项的输出结果
type Result struct {
	Name      string `json:"name"`       // 检查项名称，如 "DNS 解析"
	Target    string `json:"target"`     // 检查目标，如 "example.com"
	Status    string `json:"status"`     // "✅" 通过 / "❌" 失败
	Detail    string `json:"detail"`     // 详细信息，如 "解析到 93.184.216.34"
	LatencyMs int64  `json:"latency_ms"` // 耗时（毫秒）
	Error     string `json:"error"`      // 错误信息（成功时为空）
}

// String 格式化输出单行结果
func (r Result) String() string {
	if r.Error != "" {
		return fmt.Sprintf("%s %s | %s | %s", r.Status, r.Name, r.Target, r.Error)
	}
	return fmt.Sprintf("%s %s | %s | %s (%dms)", r.Status, r.Name, r.Target, r.Detail, r.LatencyMs)
}
