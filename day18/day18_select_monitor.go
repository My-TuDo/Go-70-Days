package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 创建两个模拟的“信号源”
	mysqlStatus := make(chan string)
	redisStatus := make(chan string)

	// 2. 模拟异步探测（协程在后台跑）
	go func() {
		time.Sleep(2 * time.Second)
		mysqlStatus <- "MySQL： 正常（200ms）"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		redisStatus <- "Redis： 正常（100ms）"
	}()

	fmt.Println("正在等待多源监控数据会传...")

	// 3. 【核心】使用 slect 来监听多个通道
	// 谁快谁先被执行
	for i := 0; i < 2; i++ {
		select { // 让一个协程可以等待多个通道操作
		case msg := <-mysqlStatus:
			fmt.Printf("收到报文：%s\n", msg)
		case msg := <-redisStatus:
			fmt.Printf("收到报文：%s\n", msg)
		case <-time.After(3 * time.Second): // 本质通道，超时控制，3s内都没有返回数据，就执行这个 case
			fmt.Println("等待超时")
		}
	}
	fmt.Println("监控数据已全部收到，程序结束")
}
