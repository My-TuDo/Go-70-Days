package main

import (
	"fmt"
	"math"
	"time"
)

// simulateRequest 模拟一个不稳定的网络请求
func simulateRequest(attempt int) error {
	if attempt < 3 {
		return fmt.Errorf("网络超时（模拟）")
	}
	return nil // 第三次终于成功了
}

func main() {
	maxRetries := 5

	fmt.Println("自动化运维任务启动，开启自愈重试模式...")

	for i := 1; i <= maxRetries; i++ {
		err := simulateRequest(i)
		if err == nil {
			fmt.Printf("第%d次尝试：成功！\n", i)
			break
		}

		fmt.Printf("第 %d 次尝试失败：%v\n", i, err)

		if i == maxRetries {
			fmt.Println("[严重告警]：达到最大重试次数，任务失败！")
			return
		}

		// [核心算法]：指数退避（2的n次方）
		// 计算下次重试需要等待的时间
		waitTime := time.Duration(math.Pow(2, float64(i))) * time.Second

		fmt.Printf("等待 %v 后进行下次重试...\n", waitTime)
		time.Sleep(waitTime)
	}
}
