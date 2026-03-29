package main

import (
	"fmt"
	"net"
	"time"
)

func main121() {
	address := "api-takumi.mihoyo.com:443" // 米哈游 API 服务器的 HTTPS 端口
	timeout := 3 * time.Second

	fmt.Printf("正在尝试连接到 %s...\n", address)

	// 记录开始时间
	start := time.Now()

	// 核心调用： 尝试拨号
	conn, err := net.DialTimeout("tcp", address, timeout)

	// 计算耗时
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("拨测失败： %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Printf("拨测成功！连接建立，响应时间: %v\n", duration)
}
