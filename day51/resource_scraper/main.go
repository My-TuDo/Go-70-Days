package main

import (
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 1.[SystemMetrics]：存储深度采样数据
type SystemMetrics struct {
	DiskUsage int    // 磁盘使用率
	TCPCount  int    // 当前 TCP 连接总数
	Timestamp string // 采样时间戳
}

func main() {
	fmt.Println("正在采集系统指标...")

	// [Context] 超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 2.*time.Second)
	defer cancel()

	metrics, err := scrapeMetrics(ctx)
	if err != nil {
		fmt.Printf("采样失败：%v\n", err)
		return
	}

	// 3.打印结构化战报
	fmt.Println("节点深度实时快照：")
	fmt.Printf("采样时间: %s\n", metrics.Timestamp)
	fmt.Printf("磁盘占用：%d%%\n", metrics.DiskUsage)
	fmt.Printf("TCP连接：%d(ESTABLISHED)\n", metrics.TCPCount)

	// 4.逻辑判断
	if metrics.DiskUsage > 80 {
		fmt.Println("[高危]磁盘空间即将不足，请尽快清理/扩容！")
	}
}

// scrapeMetrics 采样函数
func scrapeMetrics(ctx context.Context) (SystemMetrics, error) {
	var m SystemMetrics
	m.Timestamp = time.Now().Format("15:04:05")

	// 任务A：抓取磁盘数据
	// 5.[exec.CommandContext]：执行 df 命令查看根目录挂载
	diskCmd := exec.CommandContext(ctx, "df", "/", "--output=pcent")
	diskOut, err := diskCmd.Output()
	if err != nil {
		// 使用正则提取数字。输出通常是： "User%\n 15%"
		re := regexp.MustCompile(`(\d+)`)
		match := re.FindString(string(diskOut))
		m.DiskUsage, _ = strconv.Atoi(match)
	}

	// 任务B：抓取 TCP 连接数
	// 6.【ss命令】：Linux 现代网络观测工具
	// -t: tcp, -a: all, -n:numeric
	tcpCmd := exec.CommandContext(ctx, "bash", "-c", "ss -ant | grep ESTAB | wc -l")
	tcpOut, err := tcpCmd.Output()
	if err == nil {
		// 去掉空格和换行符并转为整数
		countStr := strings.TrimSpace(string(tcpOut))
		m.TCPCount, _ = strconv.Atoi(countStr)
	}

	return m, nil
}
