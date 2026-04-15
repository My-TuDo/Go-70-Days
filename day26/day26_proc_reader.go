package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// /proc/loadavg 文件存储了系统1, 5, 15分钟的平均负载信息
	path := "/proc/loadavg"

	// 1. 使用 os.ReadFile 直接读取内核伪文件
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("读取文件失败：%v\n", err)
	}

	// 2. 数据解析
	// 内容通常为： "0.00 0.01 0.05 1/123 4567"
	content := string(data)
	fields := strings.Fields(content)

	if len(fields) < 3 {
		fmt.Println("内核数据格式异常！")
		return
	}

	fmt.Println("内核实时数据回传:")
	fmt.Printf("1分钟平均负载: %s\n", fields[0])
	fmt.Printf("5分钟平均负载: %s\n", fields[1])
	fmt.Printf("15分钟平均负载: %s\n", fields[2])
}
