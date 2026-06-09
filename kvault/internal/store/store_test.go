// 测试

// 基础读写
func TestSetAndGet(t *testing.T) {
	s := New()
	s.Set("key1", "value1")

	val, ok := s.Get("key1")
	// 期望： ok == true, val == "value1"
	// 验证结果
	if !ok {
		t.Errorf("Expected key1 to exist")
	}
	if val != "value1" {
		t.Errorf("Expected value1, got %s", val)
	}
}

// 不存在的key
func TestGetMissing(t *testing.T) {
	s := New()
	_, ok := s.Get("not_exists")
	// 期望： ok == false
	if ok {
		t.Errorf("Expected missing key to not exist")
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
		}
	}
	wg.Wait()
}