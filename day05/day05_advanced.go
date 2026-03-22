// 有概率失败的采矿系统
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func workerAdvanced(id int, mine chan string, wg *sync.WaitGroup) {
	// 1. 工人出门前，告诉计数器：我干完了会报到
	defer wg.Done()

	fmt.Printf("工人%d号：进入矿洞...\n", id)

	// 模拟不确定的工作时间（0-2s）
	// 在Go 1.20+ 中，直接调用rand.Intn是并发安全的且已自动播种
	wrokTime := rand.Intn(3)
	time.Sleep(time.Duration(wrokTime) * time.Second)

	// 2. 模拟成功率： 50%概率挖不到
	if rand.Intn(10) < 5 {
		fmt.Printf("工人%d号：挖矿失败了，什么都没挖到...\n", id)
		return // 没挖到，直接结束
	}

	// 3. 挖到了，送入管道
	mine <- fmt.Sprintf("工人%d号：挖到了一块矿石！", id)
}

func main052() {
	// 注意：这里不再需要，Go 1.20之后会自动播种随机数生成器
	//rand.Seed(time.Now().UnixNano()) // 初始化设置随机数种子

	mine := make(chan string)
	var wg sync.WaitGroup

	workerCount := 10

	fmt.Println("矿山项目启动！")

	// 1. 启动工人协程
	for i := 1; i <= workerCount; i++ {
		wg.Add(1) // 告诉计数器：又有一个工人要干活了
		go workerAdvanced(i, mine, &wg)
	}

	// 2. [核心] 启动监控协程（Monitor Goroutine）
	// 它的唯一职责就是等所有工人干完活了，才关闭管道
	go func() {
		wg.Wait()   // 等所有工人都报到（干完活了）： 阻塞等待所有Done()
		close(mine) // 关闭管道，告诉主程序没有更多矿石了： 关闭通道，发出停止信号
		fmt.Println("监控员：所有工人都干完了，通道关闭了！")
	}()

	// 3. [核心] 主程序收货（Consumer）
	// range 会一直阻塞等待，直到管道被close且里面的数据被取完
	count := 0
	for result := range mine {
		count++
		fmt.Printf("收到矿石：%s\n", result)
	}

	fmt.Printf("任务总结：派出 %d 人，成功收到了%d块矿石。\n", workerCount, count)
}
