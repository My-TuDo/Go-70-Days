package scanner

import (
	"testing"
)

// 测试 CountLevel 函数
func TestCountLevel(t *testing.T) {

	tests := []struct {
		name string
		line string
		want LogLevel
	}{
		{"ERROR行", "2026/01/01 [ERROR] 数据库超时", LevelError},
		{"INFO行", "2026/04/23 [INFO] 服务启动", LevelInfo},
		{"WARN行", "2026/07/15 [WARN] 内存使用过高", LevelWarn},
		{"无级别行", "2026/10/30 12:00:00 无级别日志", LevelInfo}, // 默认 INFO
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLevel(tt.line); got != tt.want {
				t.Errorf("countLevel(%q) = %v，期望 %v", tt.line, got, tt.want)
			}
		})
	}
}
