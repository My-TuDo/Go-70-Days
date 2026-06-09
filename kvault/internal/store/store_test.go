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

// 删除测试
func TestDelete(t *testing.T) {
	s := New()

	tests := []struct {
		name string
		key  string
	}{
		{"删除存在的键", "key1"},
		{"删除不存在的键", "key2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.Set(tt.key, "value")

			s.Delete(tt.key)

			_, ok := s.Get(tt.key)
			if ok {
				t.Errorf("Delete(%q)后，Get仍然返回 ok=true", tt.key)
			}
		})
	}
}

// 列出所有key
func TestListKeys(t *testing.T) {
	s := New()

	keys := []string{
		"key1",
		"key2",
		"key3",
	}

	for _, key := range keys {
		s.Set(key, "value")
	}

	listedKeys := s.ListKeys()
	if len(listedKeys) != len(keys) {
		t.Errorf("ListKeys返回的键数量 %d, 期望 %d", len(listedKeys), len(keys))
	}
}
