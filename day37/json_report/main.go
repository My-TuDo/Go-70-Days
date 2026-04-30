package main

import (
	"encoding/json" // 核心包：处理JSON
	"fmt"
	"os" // 文件操作
	"time"
)

// TaskReport 代表最终生成的运维报告
type TaskReport struct {
	TaskID    string    `json:"task_id"`
	StartTime time.Time `json:"start_time"`
	Total     int       `json:"total"`
	Results   []string  `json:"results"`
}

func main() {
	report := TaskReport{
		TaskID:    "JOB-20260430-001",
		StartTime: time.Now(),
		Total:     3,
		Results:   []string{"10.0.0.1:Success", "10.0.0.2:Success", "10.0.0.3:Success"},
	}

	// [os.Create]：创建文件，准备写入报告
	// 作用：在当前目录下创建一个文件
	file, err := os.Create("last_job_report.json")
	if err != nil {
		fmt.Printf("Error creating report file: %v\n", err)
		return
	}
	defer file.Close()

	// [json.NewEncoder]：创建JSON编码器，准备将报告写入文件
	// 作用：将Go结构体转换为JSON格式，并写入文件
	encoder := json.NewEncoder(file)
	// [SetIndent]：设置JSON格式化输出，增加可读性
	// 作用：让生成的JSON文件更易读，增加缩进和换行
	encoder.SetIndent("", "  ")

	// [encoder.Encode]：将报告编码为JSON并写入文件
	if err := encoder.Encode(report); err != nil {
		fmt.Printf("Error writing report to file: %v\n", err)
		return
	}

	fmt.Println("[Success]结构化报告已存入 last_job_report.json")
}
