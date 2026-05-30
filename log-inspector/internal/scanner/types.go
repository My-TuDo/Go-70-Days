// Package scanner 负责扫描日志文件，按日志级别分类统计
package scanner

// LogLevel 表示日志级别
type LogLevel string

const (
	LevelInfo  LogLevel = "INFO"
	LevelWarn  LogLevel = "WARN"
	LevelError LogLevel = "ERROR"
)

// LogFileResult 记录单个日志文件的扫描结果
type LogFileResult struct {
	FilePath string              `json:"file_path"` // 文件路径
	Total    int                 `json:"total"`     // 总行数
	Counts   map[LogLevel]int    `json:"counts"`    // 各级别数量
	Errors   []string            `json:"errors"`    // ERROR 级别的具体内容（取前 5 条）
}

// ScanResult 是完整扫描的结果，包含所有文件
type ScanResult struct {
	ScannedAt string          `json:"scanned_at"` // 扫描时间
	Files     []LogFileResult `json:"files"`      // 每个文件的扫描结果
	Summary   Summary         `json:"summary"`    // 汇总
}

// Summary 是所有文件的汇总统计
type Summary struct {
	TotalFiles   int            `json:"total_files"`   // 扫描的文件数
	TotalLines   int            `json:"total_lines"`   // 总行数
	TotalByLevel map[LogLevel]int `json:"total_by_level"` // 按级别汇总
}
