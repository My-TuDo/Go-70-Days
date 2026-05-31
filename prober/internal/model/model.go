package model

import (
	"time"

	"gorm.io/gorm"
)

// TargetConfig 对应 config.yaml 里的单个目标配置
type TargetConfig struct {
	Name       string `yaml:"name"`
	URL        string `yaml:"url"`
	TimeoutSec int    `yaml:"timeout_sec"`
}

// ProbeConfig 对应整个 config.yaml
type ProbeConfig struct {
	Targets []TargetConfig `yaml:"targets"`
}

// ProbeResult 是每次拨测的结果记录，存入 MySQL
// GORM 会自动创建表名 probe_results
type ProbeResult struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	TargetName string `json:"target_name"`  // 目标名称，如 "baidu"
	TargetURL  string `json:"target_url"`   // 目标 URL
	StatusCode int    `json:"status_code"`  // HTTP 状态码，0 表示连接失败
	LatencyMs  int64  `json:"latency_ms"`   // 响应时间（毫秒）
	Success    bool   `json:"success"`      // 是否成功
	ErrorMsg   string `json:"error_msg"`    // 错误信息（成功时为空）
}

// ProbeCache 是存在 Redis 里的最新状态
// 不存数据库，只存 Redis，用 JSON 序列化
type ProbeCache struct {
	TargetName string `json:"target_name"`
	TargetURL  string `json:"target_url"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	LatencyMs  int64  `json:"latency_ms"`
	CheckedAt  string `json:"checked_at"` // 拨测时间
}
