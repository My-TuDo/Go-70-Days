package main

import (
	"bufio" // 作用：将原始字节流包装成“逐行读取”模式
	"fmt"
	"os/exec" // 作用：提供运行外部命令的功能
)

func main() {
	// 实时调用系统的 ping 命令，持续观察网络波动
	// 模拟执行 ping -c 5 baidu.com
	cmd := exec.Command("ping", "-c", "5", "baidu.com")

	// [cmd.StdoutPipe()]
	// 作用：获取命令的标准输出管道，以便我们可以实时读取命令的输出！！！
	// 意义：允许 go 程序在子进程运行的过程中，“偷听”它的实时输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("获取标准输出管道失败: %v\n", err)
		return
	}

	// [cmd.Start()]
	// 作用：非阻塞地启动程序
	// 区别：之前使用的 Run() 会等待程序结束才返回；Start() 是一下令就立即执行下一行。
	if err := cmd.Start(); err != nil {
		fmt.Printf("启动命令失败: %v\n", err)
		return
	}

	// [bufio.NewScanner(r)]
	// 作用：利用扫描器接住管道里的水流
	scanner := bufio.NewScanner(stdout)

	fmt.Println("正在实时同步子进程的网络探测报文...")

	// 实时处理流
	go func() {
		for scanner.Scan() {
			// 每当 Ping 命令产生一行新的输出，这里就会立即捕捉到
			line := scanner.Text()
			fmt.Printf("Ping 输出: %s\n", line)
		}
	}()

	// [cmd.Wait()]
	// 作用：主程序在此处等待，直到子进程结束
	// 逻辑：当 ping 命令完成它的 5 次探测后，子进程会自然结束，主程序才会继续往下走。
	if err := cmd.Wait(); err != nil {
		fmt.Printf("等待命令完成失败: %v\n", err)
	}

	fmt.Println("网络探测完成！")
}
