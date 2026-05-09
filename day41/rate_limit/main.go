package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate" // 限流包
)

func main() {
	// [工程环境]：模拟给 10 个服务器下补丁，但策略规定：
	// 每秒最多只能下发 2 个补丁，超过的请求需要排队等待，最多预存 3 个令牌

	// 1. 【rate.NewLimiter】
	// 参数1：limit (每秒生成的令牌数)
	// 参数2：burst (令牌桶的容量，预存的令牌数)
	limiter := rate.NewLimiter(2, 3)

	fmt.Println("自动化部署已就位，限流策略：2 req/s")

	ctx := context.Background()

	for i := 1; i <= 10; i++ {
		// 2. 【limiter.Wait(ctx)】等待获取令牌
		start := time.Now()
		err := limiter.Wait(ctx)
		if err != nil {
			fmt.Printf("等待令牌出错：%v\n", err)
			return
		}

		// 模拟下发补丁的操作
		fmt.Printf("[ %s ]正在给服务器 %d 下载补丁，下发补丁，耗时: %v\n", time.Now().Format("15:04:05"), i, time.Since(start))

		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("所有服务器补丁下发完成")
}
