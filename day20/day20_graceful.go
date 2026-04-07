package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. 准备一个接受信号的通道
	sigChan := make(chan os.Signal, 1)

	// 2. 订阅我们关心的信号，这里以 SIGINT（Ctrl+C）和 SIGTERM（终止信号）为例
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("程序正在运行，按 Ctrl+C 退出...")

	// 模拟一直在干活
	go func() {
		for {
			time.Sleep(2 * time.Second)
		}
	}()

	// 3.【核心】阻塞在这里，直到接受到信号
	sig := <-sigChan
	fmt.Printf("\n 捕获到信号： %v\n", sig)

	fmt.Println("正在保存数据到数据库，请稍后...")
	time.Sleep(2 * time.Second) // 模拟清理过程

	fmt.Println("资源已释放，程序安全退出。")
}
