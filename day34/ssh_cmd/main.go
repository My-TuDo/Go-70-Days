package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	// 1. 配置阶段
	config := &ssh.ClientConfig{
		User: "xzh",
		Auth: []ssh.AuthMethod{
			ssh.Password("1234567"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	// 2. 建立连接
	addr := "192.168.36.154:22" // 连接自己
	fmt.Printf("正在尝试 SSH 连接 %s\n", addr)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
		return
	}
	defer client.Close()

	// 3. 创建会话
	session, err := client.NewSession()
	fmt.Printf("SSH 连接成功，正在尝试创建 session...\n")
	if err != nil {
		log.Fatalf("创建会话失败: %v", err)
		return
	}
	defer session.Close()

	// 4. 执行命令
	fmt.Printf("正在执行命令: uptime\n")
	output, err := session.CombinedOutput("uptime")
	if err != nil {
		log.Fatalf("执行命令失败: %v", err)
		return
	}

	fmt.Printf("命令输出: %s\n", string(output))
}
