package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("正在执行系统巡检...")

	// 1. 定义要执行的命令： 查看系统负载(uptime)
	// exec/Command("命令名", "参数1", "参数2", ...)
	cmd := exec.Command("uptime")

	// 2. 执行并捕获输出结果
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("命令执行失败：%v\n", err)
		return
	}

	// 3. 处理结果（去掉末尾换行符）
	result := strings.TrimSpace(string(output))
	fmt.Printf("系统负载信息：%s\n", result)

	// 4. 【魔改】检查特定的服务是否在运行（比如ssh）
	// 模拟执行： ps aux | grep sshd
	checkCmd := exec.Command("pgrep", "ssh")
	if err := checkCmd.Run(); err != nil {
		fmt.Println("警告！SSH 服务未运行！")
	} else {
		fmt.Println("SSH 服务正在运行。")
	}
}
