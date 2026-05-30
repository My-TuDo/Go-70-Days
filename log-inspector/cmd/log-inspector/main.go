package main

import (
	"fmt"
	"os"

	"log-inspector/internal/reporter"
	"log-inspector/internal/scanner"
)

func main() {
	// 先直接测试 scanner 包是否能正常编译和工作
	dir := "./testdata"
	if len(os.Args) >= 2 {
		dir = os.Args[1]
	}

	result, err := scanner.ScanDir(dir)
	if err != nil {
		fmt.Printf("扫描失败: %v\n", err)
		os.Exit(1)
	}

	reporter.PrintSummary(result)

	err = reporter.SaveReport(result, dir+"_report")
	if err != nil {
		fmt.Printf("保存报告失败: %v\n", err)
		os.Exit(1)
	}
}
