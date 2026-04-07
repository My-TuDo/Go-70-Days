package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 1. 模拟一个数据源
	source := strings.NewReader("Hello, World!")

	// 2. 模拟一个目的地：创建一个本地文件
	// os.Create 返回的是 *os.File ，它完美实现了 io.Writer 接口，可以用来写入数据。
	destFile, err := os.Create("system_log.txt")
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer destFile.Close() // 确保文件在函数结束时被关闭

	// 3. 【核心】使用 io.Copy 将数据从 source 复制到 destFile
	// 它是连接了 source 和 destFile 之间的桥梁，负责从 source 读取数据并写入 destFile。底层会自动分片读取，不占内存
	written, err := io.Copy(destFile, source)
	if err != nil {
		fmt.Println("复制数据失败:", err)
		return
	}
	fmt.Printf("成功复制了 %d 字节的数据到 system_log.txt\n", written)
}
