package main

// 用net包搭建一个最基础的 Echo Server （回声服务器）

/*
* net.Listen: 在服务器上开一个“收发室”，守住一个端口
* Accept(): 阻塞等待。当一个玩家(客户端)连接进来时，收发室派出一个专人（新建一个conn）去接待
* Goroutine 结合：每来一个连接，就开一个协程处理，互不干扰
 */

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close() // 处理完毕后关闭连接
	remoteAddr := conn.RemoteAddr().String()
	fmt.Printf("新连接来自: %s\n", remoteAddr)

	// 使用 bufio 读取客户端发来的每一行数据
	reader := bufio.NewReader(conn)
	for {
		// 读取直到遇到换行符
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("连接 %s 断开: %v\n", remoteAddr, err)
			break
		}
		fmt.Printf("收到来自 %s 的消息: %s", remoteAddr, message)

		// 回复客户端（Echo）
		reply := "已收到你的报文：" + strings.ToUpper(message)
		conn.Write([]byte(reply))
	}
}

func main() {
	// 1. 监听本地 9000 端口
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Printf("监听端口失败: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP Echo Server 已启动，监听端口 9000...")

	for {
		// 2. 阻塞等待连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("接受连接失败: %v\n", err)
			continue
		}

		// 3. 为每个连接开启协程，实现高并发处理
		go handleConnection(conn)
	}
}
