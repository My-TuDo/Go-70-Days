package main

import (
	"fmt"
)

func processConfig(data interface{}) {
	// 尝试断言为string
	if val, ok := data.(string); ok {
		fmt.Printf("识别到文本配置：%s\n", val)
		return
	}

	// 尝试断言为int
	if val, ok := data.(int); ok {
		fmt.Printf("识别到数值配置：%d\n", val)
		return
	}

	// 如果都不是，输出未知类型
	fmt.Printf("未知配置类型：%T\n", data)
}

func main141() {
	fmt.Println("正在解析系统动态参数...")
	processConfig("mihoyo-server-01")
	processConfig(8080)
	processConfig(3.14) // 这个类型没有被处理，会输出未知类型
}
