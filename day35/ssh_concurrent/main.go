package main

import (
	"fmt"
	"log"
	"sync" // 👈 引入 sync 包来使用 WaitGroup
	"time"

	"golang.org/x/crypto/ssh"
)

// Task 结构体定义了一个任务，包含一个 ID 和一个 IP 地址
type Task struct {
	ID int
	IP string
}

func main() {
	// 模拟五个不同的服务器或是同一个服务器的不同任务
	ips := []string{"127.0.0.1:22", "127.0.0.1:22", "127.0.0.1:22", "127.0.0.1:22", "127.0.0.1:22"}
	command := "uptime"

	// 配置 SSH 凭证
	config := &ssh.ClientConfig{
		User:            "xzh",
		Auth:            []ssh.AuthMethod{ssh.Password("1234567")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         3 * time.Second,
	}

	// 管理并发进程
	var wg sync.WaitGroup
	for i, ip := range ips {
		wg.Add(1)

		// 启动并发进程
		go func(targetIP string, taskID int) {
			defer wg.Done()
			fmt.Printf("任务 %d: 正在连接到 %s...\n", taskID, targetIP)
			executeRemoteCommand(targetIP, config, command)
		}(ip, i+1)
	}

	wg.Wait()
	fmt.Println("所有任务已完成！")
}

func executeRemoteCommand(ip string, config *ssh.ClientConfig, cmd string) {
	// 建立连接
	client, err := ssh.Dial("tcp", ip, config)
	if err != nil {
		log.Fatalf("连接到 %s 失败: %v", ip, err)
		return
	}
	defer client.Close()

	// 创建会话
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("创建会话失败: %v", err)
		return
	}
	defer session.Close()

	// 执行命令并获取输出
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		log.Fatalf("执行命令失败: %v", err)
		return
	}

	fmt.Printf("来自 %s 的输出:\n%s\n", ip, string(output))
}
