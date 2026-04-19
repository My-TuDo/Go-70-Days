package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 定义一每 2 秒响一次的闹钟
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// 2. 定义一个 10 秒后自毁的定义器（模拟系统维护）
	done := time.After(11*time.Second)

	fmt.Println("系统心跳监测已启动...")
	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("[%v] 上报数据：CPU 2。5%%， 内存使用率 45%%， 状态：OK\n", t.Format("15:04:05"))
		case <-done:
			fmt.Println("系统维护中，心跳监测已停止。")
			return
		}
	}
}
