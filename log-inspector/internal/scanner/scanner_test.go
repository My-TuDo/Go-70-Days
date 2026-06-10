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

// 测试 CalcSummary 函数
func TestCalcSummary(t *testing.T) {
	files := []LogFileResult{
		{
			FilePath: "app.log",
			Total:    100,
			Counts: map[LogLevel]int{
				LevelInfo:  80,
				LevelWarn:  15,
				LevelError: 5,
			},
		},
		{
			FilePath: "db.log",
			Total:    50,
			Counts: map[LogLevel]int{
				LevelInfo:  30,
				LevelWarn:  10,
				LevelError: 10,
			},
		},
	}

	summary := calcSummary(files)
	// 验证汇总结果
	// 期望 TotalFiles = 2
	if summary.TotalFiles != 2 {
		t.Errorf("TotalFiles = %d，期望 2", summary.TotalFiles)
	}
	// 期望 TotalLines = 150
	if summary.TotalLines != 150 {
		t.Errorf("TotalLines = %d，期望 150", summary.TotalLines)
	}
	// 期望 TotalByLevel[INFO] = 110，TotalByLevel[WARN] = 25，TotalByLevel[ERROR] = 15
	if summary.TotalByLevel[LevelInfo] != 110 {
		t.Errorf("TotalByLevel[INFO] = %d，期望 110", summary.TotalByLevel[LevelInfo])
	}
	if summary.TotalByLevel[LevelWarn] != 25 {
		t.Errorf("TotalByLevel[WARN] = %d，期望 25", summary.TotalByLevel[LevelWarn])
	}
	if summary.TotalByLevel[LevelError] != 15 {
		t.Errorf("TotalByLevel[ERROR] = %d，期望 15", summary.TotalByLevel[LevelError])
	}
}
