// 测试
package store

import (
	"fmt"
	"sync"
	"testing"
)

// 基础读写
func TestSetAndGet(t *testing.T) {
	s := New()

	tests := []struct {
		name  string // 测试用例名称
		key   string
		value string
	}{
		{"普通键值", "key1", "value1"},
		{"空值", "key2", ""},
		{"特殊字符", "key3", "!@#$%^&*()"},
		{"中文值", "key4", "你好，世界"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.Set(tt.key, tt.value)

			val, ok := s.Get(tt.key)
			if !ok {
				t.Errorf("Get(%q)返回 ok=false", tt.key)
			}
			if val != tt.value {
				t.Errorf("Get(%q) = %q, 期望 %q", tt.key, val, tt.value)
			}
		})
	}
}

// 不存在的key
func TestGetMissing(t *testing.T) {
	s := New()
	_, ok := s.Get("not_exists")
	// 期望： ok == false
	if ok {
		t.Errorf("Get(%q)返回 ok=true，期望 false", "not_exists")
	}
}

// 并发安全测试:10个进程同时读写，跑完没有 data race 算通过
func TestConcurrentSafety(t *testing.T) {
	s := New()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			s.Set(key, fmt.Sprintf("value%d", i))
			val, ok := s.Get(key)
			if !ok || val != fmt.Sprintf("value%d", i) {
				t.Errorf("Concurrent safety test failed for %s", key)
			}
		}(i)
	}
	wg.Wait()
}
