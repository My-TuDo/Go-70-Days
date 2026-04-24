package main

import (
	"context"
	"fmt"
	"net" // 核心包：提供所有底层网络连接功能
	"time"
)

func main() {
	// 模拟拨测全球加速节点的 443 端口
	targets := "api-takumi.mihoyo.com:443"

	// net.Dialer
	// 作用：这是一个定制化的拨号器，比默认的 net.Dial 更加灵活，可以设置超时、控制本地地址、使用特定的网络协议等。
	dialer := &net.Dialer{
		// [Timeout]: 物理建立连接的最大等待时间
		Timeout: 5 * time.Second,

		// [KeepAlive]: 每隔 30 秒自动发送一个微小的探测包，检查对方是否还活着
		// 意义： 防止网络防火墙因为连接长时间没数据传输而偷偷断开
		KeepAlive: 30 * time.Second,
	}

	fmt.Printf("正在建立高可靠的 TCP 连接 %s...\n", targets)

	// [dialer.DialContext]
	// 作用：在给定的上下文约束下，使用定制化的拨号器建立连接
	/*
	* 参数说明：
	* - context.Background(): 这是一个空的上下文，表示没有特定的截止时间或取消信号。它是所有上下文的根。
	* - "tcp": 指定使用 TCP 协议进行连接
	* - targets: 目标地址，格式为 "host:port"
	 */
	conn, err := dialer.DialContext(context.Background(), "tcp", targets)
	if err != nil {
		fmt.Printf("连接失败: %v\n", err)
		return
	}
	defer conn.Close() // 确保连接在使用完后被关闭

	// [conn.LocalAddr()] 与 [conn.RemoteAddr()]
	// 作用：获取连接的本地和远程地址信息(IP 和端口)
	// 意义：在日志中记录数据是从哪个网卡发出去的，发到了哪个 IP
	fmt.Printf("连接成功！\n")
	fmt.Printf("本地地址: %s\n", conn.LocalAddr().String())
	fmt.Printf("远程地址: %s\n", conn.RemoteAddr().String())
}
