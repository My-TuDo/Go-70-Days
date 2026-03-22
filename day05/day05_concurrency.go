package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1. 定义工人函数
func worker(id int, mine chan string) {
	fmt.Printf("工人%d开始挖矿了\n", id)

	// 模拟挖矿耗时
	// 魔改任务：每个工人随机挖矿1-3秒，看看结果会是什么样子
	time.Sleep(time.Duration(1+rand.Intn(3)) * time.Second)

	// 挖到矿了，发送到管道
	result := fmt.Sprintf("工人%d挖到了一块矿石", id)
	mine <- result
}

func main051() {
	// 2. 创建一个"矿石转运管道"
	mine := make(chan string)

	fmt.Println("矿山项目启动！")

	// 3. 开启 3 个协程（分身），同时挖矿
	// 魔改任务：开启10个工人，看看结果会是什么样子
	for i := 1; i <= 10; i++ {
		go worker(i, mine)
	}

	// 4. 主程序在管子出口等着收货
	fmt.Println("等待矿石运出...")

	for i := 1; i <= 10; i++ {
		// 这里会发生“阻塞”，直到管子里真有东西出来
		data := <-mine
		fmt.Printf("收到矿石：%s\n", data)
	}

	// 魔改任务：统计一共收到了多少矿石，看看结果会是什么样子
	fmt.Printf("总共收到了%d块矿石\n", 10)

	fmt.Println("今日开采完毕，下班！")
}
