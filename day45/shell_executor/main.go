package main

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

// CommandResult 封装执行结果，便于 进行后续逻辑判定
type CommandResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

// RunShellCommand 执行函数
func RunShellCommand(ctx context.Context, shellCmd string) (CommandResult, error) {
	// 1.[exec.CommandContext]
	// 作用：创建一个带上下文的命令对象
	// 意义：如果脚本运行太久（超过 Context 的 Deadline），它会被强制杀死，避免资源泄露
	cmd := exec.CommandContext(ctx, "bash", "-c", shellCmd)

	// 2.[bytes.Buffer]
	// 作用：准备两个 “缓冲区” 来捕获标准输出（Stdout）和标准错误（Stderr）
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// 3.[cmd.Run()]
	// 作用：执行命令并等待完成
	err := cmd.Run()

	result := CommandResult{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
	}

	// 4.[错误处理]
	if err != nil {
		// 尝试拉取 Linux 的退出状态码（如 127 代表找不到命令， 1 代表常规错误）
		if exitError, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitError.ExitCode()
		}
		return result, err
	}
	result.ExitCode = 0
	return result, nil
}

func main() {
	// [工程环境]：模拟自动化巡检逻辑：检查磁盘空间并列出当前目录文件
	fmt.Println("自动化脚本启动中...")

	// 设定 2 秒超时，防止脚本卡死
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 模拟一段复杂的 Shell 逻辑
	myScript := `echo "--- Disk Usage ---"; df -h | head -n 2; echo " --- File List ---"; ls -l | head -n 3`

	res, err := RunShellCommand(ctx, myScript)

	if err != nil {
		fmt.Printf("脚本执行失败[Code %d]\n", res.ExitCode)
		fmt.Printf("错误输出: %s\n", res.Stderr)
		return
	}

	// 成功路径：打印标准输出
	fmt.Printf("脚本执行成功:\n%s\n", res.Stdout)
}
