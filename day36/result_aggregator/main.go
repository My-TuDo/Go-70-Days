package main

import (
	"fmt"
	"sync"
	"time"
)

// 结果收集器

// [Result结构体]：标准数据封装
// 包含目标标识、执行结果、是否报错，这样主程序才能进行统计
type ExecResult struct {
	Host    string
	Output  string
	Success bool
	Cost    time.Duration // 执行耗时
}

func mockSSHWork(host string, resultChan chan<- ExecResult, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	// 模拟执行命令
	time.Sleep(time.Duration(len(host)*100) * time.Millisecond)

	// 封装结果：不在子进程里打印，而是发回管道
	res := ExecResult{
		Host:    host,
		Output:  fmt.Sprintf("Uptime for %s: 2 days", host),
		Success: true,
		Cost:    time.Since(start),
	}
	resultChan <- res
}

func main() {
	hosts := []string{"10.0.0.1", "10.0.0.2", "10.0.0.3"}

	resultChan := make(chan ExecResult, len(hosts)) // 带缓冲的管道，避免阻塞

	var wg sync.WaitGroup

	for _, host := range hosts {
		wg.Add(1)

		go mockSSHWork(host, resultChan, &wg)
	}

	// 异步关闭通道， 主进程不会卡住
	go func() {
		wg.Wait()
		close(resultChan) // 结果收集完成，关闭管道
	}()

	// 主进程负责收集结果并统计
	for res := range resultChan {
		status := "PASS"
		if !res.Success {
			status = "FAIL"
		}
		fmt.Printf("%-10s | %-10v | %-10s\n", res.Host, res.Cost, status)
	}
	fmt.Println("所有任务已完成！")
}
