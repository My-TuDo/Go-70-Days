package main

import (
	"fmt"
	"os"

	"netdiag/internal/checker"
	"netdiag/internal/reporter"
)

func main() {
	// 用法: go run ./cmd/netdiag <hostname>
	// 示例: go run ./cmd/netdiag example.com
	//
	// 需要做的事：
	//  1. 从 os.Args 获取目标主机名
	//  2. 并行执行 4 项检查（用 goroutine + channel 收集结果）
	//  3. 用 reporter.PrintResults 打印结果

	// 检查参数
	if len(os.Args) < 2 {
		fmt.Println("用法: go run ./cmd/netdiag <主机名>")
		fmt.Println("示例: go run ./cmd/netdiag example.com")
		os.Exit(1)
	}

	// 获取目标主机名
	host := os.Args[1]

	// 并行执行检查
	results := make(chan checker.Result, 4)

	// 启动 goroutine 执行检查
	go func() { results <- checker.CheckDNS(host) }()
	go func() { results <- checker.CheckTCP(host, "443") }()
	go func() { results <- checker.CheckHTTP("https://" + host) }()
	go func() { results <- checker.CheckTLS(host) }()

	// 收集结果
	var resultList []checker.Result
	for i := 0; i < 4; i++ {
		resultList = append(resultList, <-results)
	}

	reporter.PrintResults(resultList)
}
