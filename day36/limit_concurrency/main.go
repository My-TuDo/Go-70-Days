package main

import (
	"fmt"
	"sync"
	"time"
)

// 带闸门的并发控制
func main() {
	taskCount := 10     // 共有 10 个任务
	maxConcurrency := 3 // 最多允许 3 个任务同时执行

	// 创建一个带缓冲的通道
	sem := make(chan struct{}, maxConcurrency)

	var wg sync.WaitGroup
	fmt.Println("开始执行任务...")
	for i := 1; i <= taskCount; i++ {
		wg.Add(1)
		sem <- struct{}{} // 占用一个槽位

		go func(taskID int) {
			defer wg.Done()
			defer func() { <-sem }() // 释放一个槽位

			fmt.Printf("任务 %d 开始执行\n", taskID)
			time.Sleep(2 * time.Second) // 模拟任务执行时间
			fmt.Printf("任务 %d 执行完成\n", taskID)
		}(i)
	}
	wg.Wait()
	fmt.Println("所有任务已完成！")
}
