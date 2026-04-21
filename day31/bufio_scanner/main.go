package main

import (
	"bufio" // [知识点] 缓冲 I/O 包，提供了 Scanner 类型用于逐行读取输入， 是处理大数据的核心
	"fmt"
	"os"
	"strings"
)

func main() {
	// 假设这是某台服务器上的生产日志路径
	logFilePath := "production.log"

	// 模拟创建一个日志文件
	createDummyLog(logFilePath)

	// 1. 只读模式打开日志文件
	file, err := os.Open(logFilePath)
	if err != nil {
		fmt.Printf("无法打开日志文件: %v\n", err)
		return
	}
	defer file.Close() // 确保在函数结束时关闭文件

	// 2. [核心工具] 基于文件创建一个扫描器
	// Scanner 默认缓冲区三 64KB， 它会一行一行地读，内存占用较小，适合处理大文件
	scanner := bufio.NewScanner(file)
	fmt.Println("正在扫描包含【ERROR】的日志行...")

	errorCount := 0
	for scanner.Scan() { // scanner.Scan() 让扫描器读取下一行文本，返回 true 表示成功读取到一行， false 表示没有更多行或发生错误
		// 读取当前行的文本
		line := scanner.Text()

		// 3. [业务逻辑 过滤错误日志
		if strings.Contains(line, "ERROR") {
			errorCount++
			fmt.Printf("发现故障点%d: %s\n", errorCount, line)
		}
	}

	// 4. [必] 检查扫描过程中是否出错（如硬盘坏道）
	if err := scanner.Err(); err != nil {
		fmt.Printf("扫描日志文件时发生错误: %v\n", err)
	}
	fmt.Printf("扫描完成，共发现%d条错误日志。\n", errorCount)
}

// 辅助函数：快速生成模拟日志
func createDummyLog(path string) {
	content := "INFO: User Login \nERROR: Database Timeout\nINFO: Item Picked\nERROR: Network Lag\n"
	os.WriteFile(path, []byte(content), 0644) // WriteFile 会覆盖原文件内容，如果文件不存在则创建, 0644 是文件权限
}
