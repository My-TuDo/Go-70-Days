package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 定义一每 2 秒响一次的闹钟
	// Ticker 会持有一个 Channel ，每隔固定时间往里面塞一个时间戳
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// 2. 定义一个 10 秒后自毁的定义器（模拟系统维护）
	// After 也会返回一个 Channel， 10 秒后只发送一次信号
	done := time.After(11 * time.Second)

	fmt.Println("系统心跳监测已启动...")
	// 3. 多路复用循环
	for {
		select {
		// 情况一：每当 ticker 的 Channel 收到一个时间戳，就上报一次系统状态
		case t := <-ticker.C:
			// 如果是生产环境，这里可以改成上报实际的系统状态，比如 CPU 和内存使用率
			// 然后发往 Prometheus 或是存入 MySQL
			fmt.Printf("[%v] 上报数据：CPU 2。5%%， 内存使用率 45%%， 状态：OK\n", t.Format("15:04:05"))
		// 情况二：当 done 的 Channel 收到信号，说明系统维护时间到了，停止心跳监测
		case <-done:
			fmt.Println("系统维护中，心跳监测已停止。")
			return
		}
	}
}
