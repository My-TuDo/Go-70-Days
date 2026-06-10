package scanner

import (
	"os"
	"path/filepath"
	"strings"
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

// 测试 ScanFile 函数
func TestScanFile(t *testing.T) {
	// 创造一个临时日志文件
	content := "2026/01/01 [INFO] 服务启动\n2026/01/01 [WARN] 内存使用过高\n2026/01/01 [ERROR] 数据库超时"
	tmpFile := filepath.Join(t.TempDir(), "test.log")
	os.WriteFile(tmpFile, []byte(content), 0644)

	result, err := scanFile(tmpFile)
	if err != nil {
		t.Fatalf("ScanFile 返回错误: %v", err)
	}

	// 验证结果
	if result.FilePath != tmpFile {
		t.Errorf("FilePath = %q，期望 %q", result.FilePath, tmpFile)
	}
	if result.Total != 3 {
		t.Errorf("Total = %d，期望 3", result.Total)
	}
	if result.Counts[LevelInfo] != 1 {
		t.Errorf("Counts[INFO] = %d，期望 1", result.Counts[LevelInfo])
	}
	if result.Counts[LevelWarn] != 1 {
		t.Errorf("Counts[WARN] = %d，期望 1", result.Counts[LevelWarn])
	}
	if result.Counts[LevelError] != 1 {
		t.Errorf("Counts[ERROR] = %d，期望 1", result.Counts[LevelError])
	}
	if len(result.Errors) != 1 || !strings.Contains(result.Errors[0], "数据库超时") {
		t.Errorf("Errors = %v，期望包含 '数据库超时'", result.Errors)
	}

}

// 测试 ScanDir 函数
func TestScanDir(t *testing.T) {
	// 创建临时目录
	dir := t.TempDir()

	// 创建一个 .log 文件
	logContent := "[ERROR] 数据库超时\n[INFO] 服务启动"
	os.WriteFile(filepath.Join(dir, "sever.log"), []byte(logContent), 0644)

	// 创建一个非 .log 文件
	os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("Hello World"), 0644)

	result, err := ScanDir(dir)
	if err != nil {
		t.Fatalf("ScanDir 返回错误: %v", err)
	}

	// 验证结果
	if result.Summary.TotalFiles != 1 {
		t.Errorf("TotalFiles = %d，期望 1", result.Summary.TotalFiles)
	}
	if result.Summary.TotalLines != 2 {
		t.Errorf("TotalLines = %d，期望 2", result.Summary.TotalLines)
	}
	if result.Summary.TotalByLevel[LevelError] != 1 {
		t.Errorf("TotalByLevel[ERROR] = %d，期望 1", result.Summary.TotalByLevel[LevelError])
	}
	if result.Summary.TotalByLevel[LevelInfo] != 1 {
		t.Errorf("TotalByLevel[INFO] = %d，期望 1", result.Summary.TotalByLevel[LevelInfo])
	}
	if len(result.Files) != 1 || result.Files[0].FilePath != filepath.Join(dir, "sever.log") {
		t.Errorf("Files = %v，期望包含 sever.log", result.Files)
	}

}
