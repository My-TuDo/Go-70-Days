package main

import (
	"fmt"
	"sync"
)

// 互斥锁：用于保护共享资源，确保同一时间只有一个线程访问资源

// SafeCounter 一个安全的计数器结构体
type SafeCounter struct {
	mu    sync.Mutex // 互斥锁
	count int
}

// Inc 增加计数
func (c *SafeCounter) Inc() {
	c.mu.Lock()         // 加锁
	defer c.mu.Unlock() // 保证函数退出时解锁
	c.count++
}

// Value 获取当前值
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main221() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	fmt.Println("并发压力测试启动：1000个协程同时上报错误...")

	// 启动1000个协程来增加计数
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Printf("最终统计结果：%d(预期为1000)\n", counter.Value())
}
