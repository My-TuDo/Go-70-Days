package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

// 颜色常量定义（ANSI Escape Codes）
// 通过颜色快速传递状态（绿 - 正常 | 黄 - 警告 | 红 - 危险）
const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorRed    = "\033[31m"
	ColorCyan   = "\033[36m"
)

func main() {
	// 假设这是从采样器中拿到的多节点数据
	metricsList := []struct {
		Node  string
		Disk  int
		Tcp   int
		State string
	}{
		{"sh-game-srv-01", 45, 120, "HEALTHY"},
		{"bj-auth-srv-05", 85, 2500, "WARNING"},
		{"hk-proxy-srv-02", 92, 450, "CRITICAL"},
	}

	fmt.Println(ColorCyan + "=== 集群健康状态总览 ===" + ColorReset)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("")

	// 1.[tabwriter.NewWriter]：创建一个自动对其的写入器
	// 参数说明：目标（os.Stdout)、最小列宽（0）、槽位宽度（8）、填充字符（' '）、标志（0）
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)

	// 2.写入表头
	// 使用 \t （制表符）来分隔
	fmt.Fprintln(w, "NODE_NAME\tDISK_USER\tTCP_CONN\tHEALTH_STATUS")
	fmt.Fprintln(w, "---------\t---------\t--------\t-------------")

	// 3.遍历数据并染色
	for _, m := range metricsList {
		color := ColorGreen
		if m.Disk > 90 || m.State == "CRITICAL" {
			color = ColorRed
		} else if m.Disk > 80 {
			color = ColorYellow
		}

		// 格式化每一行数据
		row := fmt.Sprintf("%s\t%d%%\t%d\t%s%s%s", m.Node, m.Disk, m.Tcp, color, m.State, ColorReset)
		fmt.Fprintln(w, row)
	}

	// 4[w.Flush()]：将缓冲区内容写入输出
	w.Flush()
}
