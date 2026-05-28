package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Process Reaper - 进程收割者")

	// 1.创造一个可取消的Context：设定任务的硬性截止时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // 确保在main函数结束时取消Context

	// 2.模拟一个子进程：循环打印并睡眠的 shell，它会持续运行直到被取消
	cmd := exec.CommandContext(ctx, "bash", "-c", "for i in {1..100}; do echo 'Child running...'; sleep 1; done")

	// 设置进程组ID（PGID）：将子进程及其产生的孙子进程全部划入一个进程组
	// 防止子进程逃逸，确保我们能统一管理它们
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// 4.异步信号监听
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		fmt.Printf("\n收到信号[%v]，正在清理所有子进程...\n", sig)
		// 5.向进程组发送负数的 PID
		// 在 Linux 中，向 -PID 发送信号代表发送给该组内的所有人
		if cmd.Process != nil {
			syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		}
		os.Exit(1)
	}()

	// 执行并获取实时输出
	fmt.Println("子进程已发射，PID：", os.Getpid())
	err := cmd.Start()
	if err != nil {
		fmt.Printf("启动失败:%v\n", err)
		return
	}

	// 等待结果
	err = cmd.Wait()
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("[Timeout]任务超时， 系统已自动强杀子进程及其所有衍生进程。")
	} else if err != nil {
		fmt.Printf("任务异常终止:%v\n", err)
	} else {
		fmt.Println("任务完美执行结束。")
	}
}
