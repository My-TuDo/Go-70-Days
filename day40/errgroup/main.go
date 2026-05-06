package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
*	errgroup:是Go语言中一个非常有用的工具，用于管理一组并发执行的任务，并且能够捕获其中任何一个任务的错误。
* 	当你需要同时执行多个任务，并且希望在其中任何一个任务失败时能够及时处理错误，errgroup是一个很好的选择。
 */

// 模拟一个运维任务
func checkServer(ctx context.Context, id int) error {
	select {
	case <-ctx.Done(): // 收到取消信号
		fmt.Printf("任务 %d：收到撤退指令，停止拨测\n", id)
		return ctx.Err()
	case <-time.After(time.Duration(id) * time.Second): // 模拟耗时
		if id == 2 {
			// 故意让任务2失败，触发错误
			return fmt.Errorf("任务 %d：服务器异常！", id)
		}
		fmt.Printf("任务 %d：巡检完成\n", id)
		return nil
	}
}

func main() {
	// 1. [errgroup.WithContext]
	// 作用：创建一个组， 并返回一个带取消功能的 Context
	// 组里任何一个协程返回 error， ctx就会立刻触发取消，通知其他协程停止工作

	group, ctx := errgroup.WithContext(context.Background())

	fmt.Println("并发巡检开始..")

	for i := 1; i <= 5; i++ {
		tempID := i // 闭包陷阱处理
		// 2. [group.Go]：启动一个新的协程去执行任务
		// 代替go func() ，它会自动帮你管理 WaitGroup 的 Add 和 Done
		group.Go(func() error {
			return checkServer(ctx, tempID)
		})
	}

	// 3. [group.Wait]：等待所有协程完成，并捕获第一个错误
	if err := group.Wait(); err != nil {
		fmt.Printf("捕获到错误：%v\n", err)
	} else {
		fmt.Println("巡检通过！")
	}
}
