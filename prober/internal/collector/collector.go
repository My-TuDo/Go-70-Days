package collector

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"prober/internal/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 你需要用到的知识点：
//
// goroutine — 并发执行（参考 day05）
// channel + WaitGroup — 并发控制（参考 day17 worker pool）
// net/http.Client — HTTP 请求（参考 day33）
// GORM — 写入 MySQL（参考 day07）

// ProbeOne 探测单个目标，返回 ProbeResult
//
// 你需要实现：
//  1. 创建一个 http.Client，设置超时时间为 target.TimeoutSec
//  2. 记录开始时间
//  3. 发起 GET 请求
//  4. 计算耗时（毫秒）
//  5. 根据结果填充 ProbeResult
//     - 成功：StatusCode、LatencyMs、Success=true
//     - 失败：Success=false、ErrorMsg=错误原因
func ProbeOne(target model.TargetConfig) model.ProbeResult {
	// TODO: 在这里写你的代码
	//
	// 创建 ProbeResult
	result := model.ProbeResult{
		TargetName: target.Name,
		TargetURL:  target.URL,
	}
	client := &http.Client{
		Timeout: time.Duration(target.TimeoutSec) * time.Second,
	}

	// 创建请求
	req, err := http.NewRequest("GET", target.URL, nil)
	if err != nil {
		result.Success = false
		result.ErrorMsg = fmt.Sprintf("创建请求失败: %v", err.Error())
		return result
	}

	// 创建带超时的 ctx
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(target.TimeoutSec)*time.Second)
	defer cancel()

	// 注入
	req = req.WithContext(ctx)

	// 记录开始时间
	start := time.Now()
	resp, err := client.Do(req)
	result.LatencyMs = time.Since(start).Milliseconds()

	// 处理结果，请求失败
	if err != nil {
		result.Success = false
		result.ErrorMsg = fmt.Sprintf("请求失败: %v", err.Error())
		return result
	}
	defer resp.Body.Close()

	// 请求成功
	result.Success = true
	result.StatusCode = resp.StatusCode
	return result

}

// ProbeAll 并发探测所有目标，控制最大并发数
//
// 你需要实现：
//  1. 创建一个带缓冲 channel 做信号量，控制同时最多跑 5 个 goroutine
//  2. 遍历所有 target，每个 target 启动一个 goroutine
//  3. 每个 goroutine 执行 ProbeOne，结果发到 results channel
//  4. 等待所有 goroutine 完成
//  5. 收集所有结果返回
//
// 提示：
//   - 信号量模式参考 day36/limit_concurrency
//   - 用 sync.WaitGroup 等待所有 goroutine 完成
func ProbeAll(targets []model.TargetConfig) []model.ProbeResult {
	maxConcurrency := 5

	// 信号量：最多 maxConcurrency 个 goroutine 同时执行
	sem := make(chan struct{}, maxConcurrency) // struct{} 空结构体类型，占位，不占内存
	// 结果 channcel ，缓冲大小为目标数量，避免 goroutine 阻塞
	results := make(chan model.ProbeResult, len(targets))
	// WaitGroup 等待所有 goroutine 完成
	var wg sync.WaitGroup

	for _, target := range targets {
		wg.Add(1)
		sem <- struct{}{} // 获取一个信号量（struct{}{}为struct{}类型的具体值），超过 maxConcurrency 的 goroutine 会在这里阻塞

		go func(t model.TargetConfig) {
			defer wg.Done()
			defer func() { <-sem }()

			results <- ProbeOne(t)
		}(target)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var resultList []model.ProbeResult
	for r := range results {
		resultList = append(resultList, r)
	}
	return resultList

}

// RunProbeCycle 执行一轮完整的探测流程：
// 并发探测 → 写入 MySQL → 打印结果
//
// 你需要实现：
//  1. 调用 ProbeAll 并发探测所有目标
//  2. 遍历每个结果，用 db.Create 写入 MySQL
//  3. 打印每个目标的探测结果到终端
func RunProbeCycle(targets []model.TargetConfig, db *gorm.DB, rdb *redis.Client) {
	// TODO: 在这里写你的代码
	//

	fmt.Printf("\n=== 开始拨测 %d 个目标 ===\n", len(targets))
	results := ProbeAll(targets)

	for _, r := range results {
		// 写入 MySQL
		db.Create(&r)

		// 写入 Redis 缓存
		if rdb != nil {
			UpdateCache(rdb, r)
		}

		// 打印结果
		status := "✅"
		if !r.Success {
			status = "❌"
		}
		fmt.Printf("%s %-10s | 状态码: %d | 延迟: %dms | %s\n",
			status, r.TargetName, r.StatusCode, r.LatencyMs, r.ErrorMsg)
	}

}
