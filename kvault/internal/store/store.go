package store

import "sync"

type Store struct {
	mu   sync.RWMutex      // 写读锁：读可以并发，写互斥
	data map[string]string // 数据存储：键值对
}

// New 创建一个新的 Store 实例，初始化 map
func New() *Store {
	return &Store{data: make(map[string]string)}
}

// 写：加写锁，存 map
func (s *Store) Set(key, value string) {
	s.mu.Lock()         // 加写锁，独占访问
	defer s.mu.Unlock() // 解锁，确保函数结束时释放锁
	s.data[key] = value // 存储键值对
}

// 读：加读锁，从 map 取
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()         // 加读锁，允许多个读者并发访问
	defer s.mu.RUnlock() // 解锁，确保函数结束时释放锁
	value, ok := s.data[key]
	return value, ok
}

// 删：加写锁，删 map
func (s *Store) Delete(key string) {
	s.mu.Lock()         // 加写锁，独占访问
	defer s.mu.Unlock() // 解锁，确保函数结束时释放锁
	delete(s.data, key) // 删除键值对
}

// 遍历：加读锁， 收集所有 key
func (s *Store) ListKeys() []string {
	s.mu.RLock()          // 加读锁，允许多个读者并发访问
	defer s.mu.RUnlock() // 解锁，确保函数结束时释放锁
	keys := make([]string, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}
	return keys
}
