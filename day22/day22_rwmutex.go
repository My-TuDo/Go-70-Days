package main

import (
	"fmt"
	"sync"
)

type SafeCounter_ struct {
	mu    sync.RWMutex // 读写锁
	count int
}

// Inc 增加计数
func (c *SafeCounter_) Inc_() {
	c.mu.Lock()         // 写锁：同时只能有一个线程持有
	defer c.mu.Unlock() // 保证函数退出时解锁
	c.count++
}

// Value 获取当前值
func (c *SafeCounter_) Value_() int {
	c.mu.RLock() // 读锁：允许多个线程同时持有
	defer c.mu.RUnlock()
	return c.count
}

func main() {
	counter := SafeCounter_{}
	var wg sync.WaitGroup

	fmt.Println(" 并发压力测试启动： 1000个协程同时上报错误...")

	// 启动1000个协程来增加计数
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc_()
		}()
	}

	wg.Wait()
	fmt.Printf("最终统计结果：%d(预期为1000)\n", counter.Value_())
}
