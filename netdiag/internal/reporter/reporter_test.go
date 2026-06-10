package reporter

import (
	"netdiag/internal/checker"
	"strings"
	"testing"
)

// 测试 SaveJSON 函数
func TestSaveJSON(t *testing.T) {
	result := []checker.Result{
		{Name: "DNS 解析", Target: "example.com", Status: "成功", Detail: "93.184.216.34", LatencyMs: 12},
	}

	jsonStr, err := SaveJSON(result)
	if err != nil {
		t.Fatalf("SaveJSON 返回错误：%v", err)
	}

	// 期望：jsonStr 是合法的 JSON 字符串
	// 期望：包含 “DNS 解析”
	// 期望：包含“93.184.216.34”
	if len(jsonStr) == 0 {
		t.Errorf("SaveJSON 返回的 JSON 字符串为空")
	}
	if !strings.Contains(jsonStr, "DNS 解析") {
		t.Errorf("SaveJSON 返回的 JSON 字符串不包含 \"DNS 解析\"")
	}
	if !strings.Contains(jsonStr, "93.184.216.34") {
		t.Errorf("SaveJSON 返回的 JSON 字符串不包含 \"93.184.216.34\"")
	}
}
