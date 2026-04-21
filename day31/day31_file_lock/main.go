package main

import (
	"fmt"
	"os"
	"syscall" // [知识点] 直接调用linux系统底层接口
	"time"
)

func main() {
	lockFile := "backup_task.lock"

	// 1. 创建或打开锁文件
	f, err := os.Create(lockFile)
	if err != nil {
		fmt.Printf("无法创建锁文件: %v\n", err)
		return
	}
	defer f.Close()

	// 2. [核心] 调用系统底层 flock 尝试加锁
	// syscall.LOCK_EX: 独占锁，其他进程无法获得锁
	// syscall.LOCK_NB: 非阻塞模式，如果无法获得锁立即返回错误
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		fmt.Println("[警告] 另一个备份进程正在运行，本次任务自动退出。")
		return
	}

	// 3. 模拟正在执行运维任务
	fmt.Println("[进程A] 获得锁，正在执行备份任务...")
	time.Sleep(10 * time.Second) // 模拟任务执行时间

	// 4. 任务完成后自动释放锁（defer f.Close() 会关闭文件描述符，系统会自动释放锁）,也可以手动释放
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN) // 手动释放锁
	fmt.Println("[进程A] 备份任务完成，锁已释放。")
}
