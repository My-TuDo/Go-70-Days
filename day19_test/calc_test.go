package main

import "testing"

func TestAdd(t *testing.T) {
	// 1. 准备数据： 表格驱动测试（Table-driven tests）
	cases := []struct { // 匿名结构体切片，包含测试用例的名称、输入参数和预期结果
		name     string
		a, b     int
		expected int
	}{
		{"正数相加", 1, 2, 3},
		{"负数相加", -1, -2, -3},
		{"正负数相加", -1, 2, 1},
		{"零相加", 0, 0, 0},
	}

	// 2. 执行测试
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { // t.Run 会为每一个用例创建一个独立的子测试。
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("FAILED: %s, 得到 %d, 期望 %d", tc.name, result, tc.expected) // 不会中断程序，会记录错误并且继续运行后续用例
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) { // 基准测试：衡量代码性能的函数，函数名必须以 Benchmark 开头，并且接受一个 *testing.B 类型的参数。
	// b.N 是 Go 测试框架自动设置的循环次数，通常是一个很大的数字，用于评估函数的性能。
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}
