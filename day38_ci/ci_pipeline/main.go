package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// 定义流水线最终的战报
type PipelineReport struct {
	Timestamp time.Time `json:"timestamp"`
	TestPass  bool      `json:"test_pass"`
	BuildPass bool      `json:"build_pass"`
	Message   string    `json:"message"`
}

func main() {
	fmt.Println("CI 系统启动中...")
	report := PipelineReport{Timestamp: time.Now()}

	// 测试阶段
	fmt.Println("执行单元测试...")

	testCmd := exec.Command("go", "test", "-v", "/home/xzh/Code/github/Go-70-Days/day19_test") // 指定测试文件路径

	testCmd.Stdout = os.Stdout // 将测试输出直接显示在控制台
	testCmd.Stderr = os.Stderr // 将错误输出直接显示在控制台
	// Run()会阻塞直到完成
	if err := testCmd.Run(); err != nil {
		report.TestPass = false
		report.Message = "单元测试失败"
		saveReport(report)
		fmt.Println(report.Message)
		return // Fail fast: 测试失败直接退出，不继续构建
	}

	report.TestPass = true
	fmt.Println("单元测试通过，开始构建...")

	// 自动化编译
	fmt.Println("编译二进制文件...")

	buildCmd := exec.Command("go", "build", "-o", "sentry_app", "/home/xzh/Code/github/Go-70-Days/day34/ssh_cmd/main.go")

	if err := buildCmd.Run(); err != nil {
		report.BuildPass = false
		report.Message = "编译失败"
		saveReport(report)
		fmt.Println(report)
		return
	}

	report.BuildPass = true
	report.Message = "CI 流水线执行成功"
	// 生成报告结果
	saveReport(report)
	fmt.Println("CI流程执行完成，报告已生成")
}

// 辅助函数：将结果保存为 JSON文件
func saveReport(r PipelineReport) {
	file, _ := os.Create("ci_result.json")
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(r)
}
