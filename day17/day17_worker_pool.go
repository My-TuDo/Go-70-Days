package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()           // 函数结束时通知 WaitGroup ，计数器减一
	for task := range tasks { // 阻塞监听: 从通道里不断取任务，直到通道关闭且数据取完
		fmt.Printf("工人 %d 号正在处理任务 %d\n", id, task)
		time.Sleep(1*time.Second + time.Duration(id)*time.Second) // 模拟不同工人处理速度不同
	}
}
func main171() {
	taskCount := 10  // 总任务数
	workerCount := 3 // 限制只有三个并发

	// 创建一个带缓冲的通道， 容量为10， 用来存放任务
	tasks := make(chan int, taskCount)
	var wg sync.WaitGroup // 声明一个 WaitGroup 用来等待所有工人完成任务

	// 启动固定数量的工人
	for i := 1; i <= workerCount; i++ {
		wg.Add(1) // 每启动一个工人，计数器加一
		go worker(i, tasks, &wg)
	}

	// 塞入任务
	for i := 1; i <= taskCount; i++ {
		tasks <- i // 往通道里塞任务 ID
	}
	close(tasks) // 关闭通道，表示没有更多任务了

	// 等待所有工人完成
	wg.Wait()
	fmt.Println("所有任务已完成！")
}
