package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 模拟网络网关，监听9000端口
	// 使用底层的 Resolve + ListenTCP 以便获得对 Socket 的控制
	addr, _ := net.ResolveTCPAddr("tcp", ":9000")
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Printf("监听失败：%v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP网关已启动，监听9000端口...")

	for {
		// 阻塞等待新玩家连接
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}

		// 开启进程处理，实现高并发
		go handlePlayer(conn)
	}
}

func handlePlayer(conn *net.TCPConn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	fmt.Printf("玩家连接：%s\n", remoteAddr)

	// 设置保活探测（KeepAlive）：自动在底层发心跳包
	// 检测玩家是否因为异常断电、断网而变成“僵尸连接”
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(30 * time.Second)

	for {
		// 设置读取超时（Read Deadline）：如果玩家在接下来的 10 秒内一个字都不说，服务器就主动断开连接
		// 防御“慢连接攻击”。防止死连接长期霸占服务器的文件描述符（FD）
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			// 现象得出结果：判断是否因为超时导致的错误
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				fmt.Printf("p[超时]玩家 %s 长期无数据，系统强制断开已回收资源。\n", remoteAddr)
			} else {
				fmt.Printf("[断开]玩家 %s 已正常退出。\n", remoteAddr)
			}
			return
		}

		fmt.Printf("[%s]发来报文：%s", remoteAddr, string(buf[:n]))
		conn.Write([]byte("服务器已收到您的消息！连接有效期延长10秒...\n"))
	}
}
