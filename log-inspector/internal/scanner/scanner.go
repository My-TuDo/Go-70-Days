package scanner

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 提示：你需要用到的标准库
//
// filepath.Walk — 遍历目录（参考 day46/file_walker）
// os.Open — 打开文件
// bufio.NewScanner — 逐行读取（参考 day31/bufio_scanner）
// strings.HasPrefix / strings.Count — 字符串判断
// time.Now().Format — 生成时间戳

// ScanDir 扫描 dir 目录下所有 .log 文件，返回汇总结果
//
// 你需要实现：
//  1. 创建一个 ScanResult，把 ScannedAt 设为当前时间
//  2. 用 filepath.Walk 遍历 dir 目录
//  3. 只处理后缀为 .log 的文件（跳过目录和其他文件）
//  4. 每个 .log 文件调用 scanFile 处理
//  5. 把 scanFile 的结果追加到 ScanResult.Files
//  6. 最后计算 Summary（累加每个文件的统计值）
//  7. 返回 ScanResult
func ScanDir(dir string) (ScanResult, error) {
	result := ScanResult{
		ScannedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("访问路径时发生错误: %s\n", err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".log") {
			return nil
		}

		fileResult, err := scanFile(path)
		if err != nil {
			fmt.Printf("处理文件时发生错误: %s\n", err)
			return err
		}
		result.Files = append(result.Files, fileResult)
		return nil
	})
	if err != nil {
		fmt.Printf("遍历目录时发生错误: %s\n", err)
		return result, err
	}
	result.Summary = calcSummary(result.Files)

	return result, nil
}

// scanFile 扫描单个日志文件，统计各级别行数
//
// 你需要实现：
//  1. 用 os.Open 打开文件
//  2. 用 bufio.NewScanner 逐行读取（别忘了 defer 关闭文件）
//  3. 每行判断包含 [INFO]、[WARN]、[ERROR] 哪个关键字
//  4. 统计总行数和各级别出现次数
//  5. 如果是 ERROR 级别，保存该行内容到 Errors 切片（最多存 5 条）
//  6. 返回 LogFileResult
//
// 提示：判断级别可以用 strings.Contains(line, "[ERROR]")
func scanFile(path string) (LogFileResult, error) {
	// 只读形式打开文件。
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("无法打开文件: %s\n", err)
		return LogFileResult{}, err
	}
	defer file.Close()

	// 基于文件创建一个扫描器
	scanner := bufio.NewScanner(file)

	fmt.Printf("正在扫描文件：%s\n", path)

	// 逐行扫描文件内容
	result := LogFileResult{
		FilePath: path,
		Counts:   make(map[LogLevel]int),
	}

	for scanner.Scan() {
		line := scanner.Text()
		// 判断日志级别，更新统计
		level := countLevel(line)
		result.Counts[level]++
		result.Total++
		// 记录 ERROR 级别的具体内容（最多存 5 条）
		if level == LevelError && len(result.Errors) < 5 {
			result.Errors = append(result.Errors, line)
		}
	}

	return result, scanner.Err()
}

// countLevel 判断一行日志属于哪个级别
// 提示：用 strings.Contains 匹配 [INFO]、[WARN]、[ERROR]
func countLevel(line string) LogLevel {
	// TODO: 如果包含 "[ERROR]" 返回 LevelError
	//       如果包含 "[WARN]"  返回 LevelWarn
	//       如果包含 "[INFO]"  返回 LevelInfo
	//       否则返回 LevelInfo（默认）

	if strings.Contains(line, "[ERROR]") {
		return LevelError
	}
	if strings.Contains(line, "[WARN]") {
		return LevelWarn
	}
	if strings.Contains(line, "[INFO]") {
		return LevelInfo
	}

	return LevelInfo
}

// calcSummary 遍历所有文件结果，计算汇总数据
func calcSummary(files []LogFileResult) Summary {
	// TODO: 累加 TotalFiles、TotalLines、TotalByLevel
	//
	// 参考：
	// summary := Summary{
	//     TotalByLevel: make(map[LogLevel]int),
	// }
	// summary.TotalFiles = len(files)
	// for _, f := range files {
	//     summary.TotalLines += f.Total
	//     for level, count := range f.Counts {
	//         summary.TotalByLevel[level] += count
	//     }
	// }

	summary := Summary{
		TotalByLevel: make(map[LogLevel]int),
	}
	// 总文件数
	summary.TotalFiles = len(files)
	// 遍历每个文件的结果，累加总行数和各级别数量
	for _, f := range files {
		summary.TotalLines += f.Total
		for level, count := range f.Counts {
			summary.TotalByLevel[level] += count
		}
	}

	return summary
}
