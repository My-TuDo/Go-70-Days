package main

import (
	"fmt"
	"time"
)

// 打不死的监控器

// protect 是一个保护函数
func protect131() {
	// 灭火器必须安装在defer里面
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("警报！捕获到系统崩溃：%v\n", r)
			fmt.Println("正在自动修复并重启监控模块...")
		}
	}()

	fmt.Println("核心监控模块已启动...")
	time.Sleep(1 * time.Second)

	// --- 模拟一个由 ”空指针“ 引起的崩溃 ---
	var data *int
	fmt.Println("尝试读取数据：", *data) // 这里会引发 panic, 因为 data 是 nil
}

func main() {
	fmt.Println("系统监控器总线启动...")

	for i := 1; i <= 3; i++ {
		fmt.Printf("\n--- 第%d次启动尝试 ---\n", i)
		protect()                   // 即使 protect函数内部崩了， main 循环也不会停
		time.Sleep(1 * time.Second) // 模拟重启间隔
	}
	fmt.Println("\n即使多次崩溃， 主程序依然平稳运行，下班！")
}
