package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 查看当前CPU核心数
	// Go 默认会使用所有的逻辑 CPU ，这就算为什么 Go 并发强的原因
	fmt.Printf("CPU 核心数: %d\n", runtime.NumCPU())

	// 模拟一个正在疯狂产生协程的业务场景
	for i := 0; i < 10; i++ {
		go func(id int) {
			// 模拟协程在干活
			time.Sleep(5 * time.Second)
		}(i)
	}

	// 每秒监控一次运行时状态
	ticker := time.NewTicker(1 * time.Second)
	stop := time.After(6 * time.Second) // 6秒后停止监控

	fmt.Println("开始监控运行时状态...")

	for {
		select {
		case <-ticker.C:
			// 获取当前协程数量
			numG := runtime.NumGoroutine()

			// 获取内存状态
			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			fmt.Printf("[%s] | 活跃协程数：%d | 堆内存占用：%.2f MB\n | GC 累积次数：%d\n",
				time.Now().Format("15:04:05"),
				numG,
				float64(m.Alloc)/1024/1024, // 将字节转为MB
				m.NumGC,
			)
		case <-stop:
			fmt.Println("停止监控，程序退出.")
			return
		}
	}
}
