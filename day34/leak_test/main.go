package main

import (
	"fmt"
	"os" // 👈 引入 os 包来获取 PID
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	// 【解决问题 2】: 让程序启动时大声喊出自己的身份证号
	pid := os.Getpid()
	fmt.Printf("🆔 我的进程 PID 是: %d\n", pid)
	fmt.Println("💡 请在另一个终端输入: lsof -p", pid, "| wc -l 来观察文件描述符增长")

	config := &ssh.ClientConfig{
		User:            "xzh",
		Auth:            []ssh.AuthMethod{ssh.Password("1234567")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	fmt.Println("💣 资源熔断实验开始...")

	for i := 1; i <= 2000; i++ {
		client, err := ssh.Dial("tcp", "127.0.0.1:22", config)
		if err != nil {
			fmt.Printf("🚨 第 %d 次连接报错: %v\n", i, err)
			break
		}

		// 【解决问题 1】: 打印 client 的内存指针地址
		// %p 代表 Pointer，能让你看到每一个连接在内存里的真实位置
		// 这样编译器就不会报错“未使用”，你也通过视觉确认了连接已建立
		if i%10 == 0 {
			fmt.Printf("📈 已堆积 %d 个链接，最新地址: %p\n", i, client)
			time.Sleep(100 * time.Millisecond)
		}
	}

	fmt.Println("⏸️  资源已占满。现在去另一个终端执行命令观察吧！")
	select {}
}
