package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type CheckResult struct {
	Host    string
	Status  string
	Latency time.Duration
	Err     error
}

func main() {
	// 1.模拟待巡检的集群节点
	hosts := []string{
		"http://127.0.0.1:8080/health",
		"https://api-takumi.mihoyo.com/health",
		"http://localhost:8081/status",
	}

	// 2.并发与收集装置
	resultChan := make(chan CheckResult, len(hosts))
	var wg sync.WaitGroup

	// 3.超时控制：5s
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("启动...")

	for _, host := range hosts {
		wg.Add(1)
		// 开启
		go probeHost(ctx, host, resultChan, &wg)
	}

	// 4.异步守候
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 5.实时汇总报表
	fmt.Printf("%-30s | %-10s | %-10s\n", "节点地址", "状态码", "延迟")
	fmt.Println("--------------------------------------------------------")
	for res := range resultChan {
		if res.Err != nil {
			fmt.Printf("%-30s | ERROR | %v\n", res.Host, res.Err)
		} else {
			fmt.Printf("%-30s | %-7d | %v\n", res.Host, res.Status, res.Latency)
		}
	}
}
func probeHost(ctx context.Context, url string, ch chan<- CheckResult, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	// 构造带 Context 的请求
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	client := &http.Client{}

	resp, err := client.Do(req)

	res := CheckResult{Host: url, Latency: time.Since(start)}
	if err != nil {
		res.Err = err
	} else {
		res.Status = fmt.Sprintf("%d", resp.StatusCode)
		resp.Body.Close()
	}

	ch <- res
}
