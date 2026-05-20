package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// [工程环境]：模拟运维场景 -- 清理 /tmp ，目录下超过 7 天的旧日志
	targetDir := "./logs_demo"

	// 准备演示环境
	os.Mkdir(targetDir, 0755)
	os.Create(filepath.Join(targetDir, "old_log.txt"))

	fmt.Println("文件清理启动...")

	// 1.[filepath.Walk]：递归遍历文件夹下的每一个文件和子文件夹
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 2.[info.IsDir()]：判断文件夹还是文件
		if info.IsDir() {
			fmt.Printf("进入目录：%s\n", path)
			return nil
		}

		// 3.[info.ModTime()]：获取我呢间最后修改时间
		age := time.Since(info.ModTime())

		fmt.Printf("检查文件：%-15s | 大小：%-5d bytes | 寿命：%v\n", info.Name(), info.Size(), age)

		// 4.[逻辑判断]：如果文件大于一个小时（演示）
		if age > 1*time.Hour {
			fmt.Printf("发现过期文件，准备清理：%s\n", path)
			// os.Remove(path) // 实际删除文件
		}

		return nil
	})
	if err != nil {
		fmt.Printf("扫描过程出错：%v\n", err)
	}

}
