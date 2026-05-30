package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"log-inspector/internal/scanner"
)

// 提示：你需要用到的标准库
//
// encoding/json.MarshalIndent — 生成带缩进的 JSON（参考 day37/json_report）
// os.Create — 创建文件
// filepath.Ext — 检查文件扩展名

// GenerateReport 接收扫描结果，返回格式化的 JSON 字符串
//
// 你需要实现：
//  1. 用 json.MarshalIndent 将 result 序列化为 JSON
//     json.MarshalIndent(result, "", "  ")
//  2. 返回 JSON 字符串
//  3. 如果序列化失败，返回空字符串和错误
func GenerateReport(result scanner.ScanResult) (string, error) {
	// TODO: 在这里写你的代码

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Printf("生成 JSON 报告时发生错误: %s\n", err)
		return "", err
	}

	return string(jsonStr), nil
}

// SaveReport 将扫描结果保存到 outputPath 文件中
//
// 你需要实现：
//  1. 调用 GenerateReport 生成 JSON
//  2. 用 os.Create 创建输出文件
//  3. 把 JSON 字符串写入文件
//  4. 打印 "报告已保存到: xxx" 到控制台
//  5. 返回错误（如果有）
//
// 提示：
//
//	输出文件路径建议加上 .json 后缀
//	可以用 filepath.Ext 检查，如果没有后缀就自动加 .json
func SaveReport(result scanner.ScanResult, outputPath string) error {
	// TODO: 在这里写你的代码

	// 生成 JSON 报告
	jsonStr, err := GenerateReport(result)
	if err != nil {
		return err
	}

	// 确保文件有 .json 后缀
	if filepath.Ext(outputPath) != ".json" {
		outputPath += ".json"
	}

	// 创建输出文件
	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("创建报告文件时发生错误: %s\n", err)
		return err
	}
	defer file.Close()

	// 写入 JSON 字符串
	_, err = file.WriteString(jsonStr)
	if err != nil {
		fmt.Printf("写入报告文件时发生错误: %s\n", err)
		return err
	}

	fmt.Printf("报告已保存到: %s\n", outputPath)
	return nil

}

// PrintSummary 在终端打印简洁的统计摘要
//
// 你需要实现：在终端打印类似下面的表格：
//
// === 日志扫描报告 ===
// 扫描时间: 2025-07-18 10:30:00
// 扫描文件: 2 个
// 总行数:   30 行
// -----------------
// INFO:  20 行
// WARN:   6 行
// ERROR:  4 行  ⚠️

func PrintSummary(result scanner.ScanResult) {
	// TODO: 用 fmt.Printf 打印格式化的摘要
	// 参考 day37 的 JSON 输出思路，但这里用文本格式
	//
	// 提示：可以用 result.Summary.TotalByLevel[scanner.LevelError] 获取 ERROR 总数

	fmt.Println("=== 日志扫描报告 ===")
	fmt.Printf("扫描时间: %s\n", result.ScannedAt)
	fmt.Printf("扫描文件: %d 个\n", len(result.Files))
	fmt.Printf("总行数:   %d 行\n", result.Summary.TotalLines)
	fmt.Println("-----------------")
	fmt.Printf("INFO:  %d 行\n", result.Summary.TotalByLevel[scanner.LevelInfo])
	fmt.Printf("WARN:   %d 行\n", result.Summary.TotalByLevel[scanner.LevelWarn])
	errorCount := result.Summary.TotalByLevel[scanner.LevelError]
	if errorCount > 0 {
		fmt.Printf("ERROR:  %d 行  ⚠️\n", errorCount)
	} else {
		fmt.Printf("ERROR:  %d 行\n", errorCount)
	}

}
