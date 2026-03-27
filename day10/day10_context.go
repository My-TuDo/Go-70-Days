package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done(): // 收到 “撤退” 信号
			fmt.Printf("%s 收到撤退信号，停止工作\n", name)
			return
		default:
			fmt.Printf("%s 正在工作...\n", name)
			time.Sleep(500 * time.Millisecond) // 模拟工作
		}
	}
}

func main101() {
	// 1. 创建一个带 “3s超时” 的context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // 结束时释放资源

	// 2. 启动协程
	go worker(ctx, "网络监控官")

	// 3. 模拟主程序忙碌
	fmt.Println("监控系统启动，限时3s...")

	// 等待5s， 看看会发生什么
	time.Sleep(5 * time.Second)
	fmt.Println("主程序结束。")
}
