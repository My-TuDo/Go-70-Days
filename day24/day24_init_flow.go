package main

import (
	"fmt"
)

var (
	// 1. 变量首先被初始化
	systemStatus = "待机..."
)

func init() {
	// 2. init 函数接着运行
	fmt.Println("[Step 1] 正在加载内核模块...")
	systemStatus = "运行中..."
}

func init() {
	// 3. 第三个 init 按照顺序运行
	fmt.Println("[Step 2] 安全防火墙已开启...")
}

func main() {
	// 4. 最后才是 main
	fmt.Printf("[Step 3] 系统状态: %s\n", systemStatus)
}
