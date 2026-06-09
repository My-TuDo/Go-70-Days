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