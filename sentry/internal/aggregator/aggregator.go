package aggregator

import (
	"fmt"
	"sync"
	"time"
)

// Alert 表示一条报警消息
type Alert struct {
	Title   string `json:"title"`   // 报警标题
	Level   string `json:"level"`   // 级别: info / warn / critical
	Source  string `json:"source"`  // 来源: prober / prometheus / ...
	Message string `json:"message"` // 详细描述
}

// Aggregator 报警聚合器，用于去重
// 同一种报警（相同 title + source）在 window 时间内只推送一次
type Aggregator struct {
	mu     sync.RWMutex
	window time.Duration        // 去重窗口，例如 5 分钟
	cache  map[string]time.Time // key → 上次推送时间
}

// New 创建一个聚合器，window 是去重窗口时长
// 例如 New(5 * time.Minute) 表示 5 分钟内重复报警不发送
func New(window time.Duration) *Aggregator {
	return &Aggregator{
		window: window,
		cache:  make(map[string]time.Time),
	}
}

// alertKey 生成报警的唯一标识，用于去重判断
// 相同 title + source 视为同一种报警
//
// 提示：用 fmt.Sprintf("%s:%s", a.Title, a.Source) 拼接
func alertKey(a Alert) string {
	// TODO: 拼接 title 和 source，返回字符串
	return fmt.Sprintf("%s:%s", a.Title, a.Source)
}

// ShouldSend 判断这条报警是否应该推送
// 规则：如果 cache 中存在相同 key，且距离上次推送时间小于 window，则不推送
//
//	否则记录当前时间到 cache，返回 true
//
// 提示：
//   - 用 alertKey(a) 生成 key
//   - 用 s.mu.RLock() 读取缓存，s.mu.Lock() 写入缓存
//   - time.Now().Before(lastTime.Add(s.window)) 判断是否在窗口内
func (s *Aggregator) ShouldSend(a Alert) bool {
	// TODO: 实现去重逻辑
	key := alertKey(a) // 生成一个 key

	s.mu.RLock() // 读锁
	lastTime, exists := s.cache[key]
	s.mu.RUnlock() // 读完立即释放读锁

	// 如果存在，并且当前时间在上次推送时间加窗口时间之前，说明在窗口内，不应该推送
	if exists && time.Now().Before(lastTime.Add(s.window)) {
		return false // 在窗口内，返回 false
	}

	// 不在窗口内，更新缓存并返回 true
	s.mu.Lock()               // 写锁
	s.cache[key] = time.Now() // 更新缓存
	s.mu.Unlock()

	return true
}
