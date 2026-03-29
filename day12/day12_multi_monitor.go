package main

import (
	"context"
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

// 1.定义站点探测结果结构体
type Result struct {
	Site    string
	Latency time.Duration
	Err     error
}

func probe(ctx context.Context, site string, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now() // 记录开始时间
	// 使用带 context 的拨号器，支持超时强断
	// 建立连接
	d := net.Dialer{} // 创建一个拨号器
	conn, err := d.DialContext(ctx, "tcp", site)

	// 判断成功，将连接释放
	if err == nil {
		conn.Close()
	}

	// 即使失败，也要把结果送回管道， 告诉主程序“我跑完了“
	// conn 只是一个检测是否能连通的工具，不需要传回主程序，所以不放在 Result 里
	ch <- Result{
		Site:    site,              // 站点地址
		Latency: time.Since(start), // 耗时
		Err:     err,               // 错误信息（如果有）
	}
}

func main123() {
	// 待探测的站点列表（米哈游部分接口及其常用 DNS)
	sites := []string{
		"api-takumi.mihoyo.com:443",
		"bbs-api.mihoyo.com:443",
		"114.114.114.114:53",
		"8.8.8.8:53",
		"google.com:443",
	}

	// 2. 初始化环境
	results := make([]Result, 0)
	resChan := make(chan Result, len(sites))
	var wg sync.WaitGroup

	// 设定整个监控任务的总超时为 5 秒
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("开启全球加速探测流...")

	// 3. 并发开发
	for _, site := range sites {
		wg.Add(1)
		go probe(ctx, site, resChan, &wg)

	}

	// 4.这里的逻辑很巧妙： 用协程去等 wg，主程序去收管道
	go func() {
		wg.Wait()
		close(resChan)
	}()

	// 5. 收集结果
	for res := range resChan {
		results = append(results, res)
	}

	// 6.【核心魔改】按延迟排序
	sort.Slice(results, func(i, j int) bool {
		return results[i].Latency < results[j].Latency
	})

	// 7. 视觉化展示
	fmt.Println("\n	--- 探测结果排序榜（从快到慢）")
	for i, r := range results {
		status := "正常"
		if r.Err != nil {
			status = "超时/失败"
		}
		fmt.Printf("[%d] 耗时： %-10v | 状态: %s | 目标： %s\n", i+1, r.Latency, status, r.Site)

	}
}
