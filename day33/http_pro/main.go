package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func main() {
	// 模拟一个需要极高稳定性的监控数据采集器
	// 不再使用 http.DefaultClient, 而是手动打造一个

	// [http.Transport]: 底层的“运输车”
	// 作用：控制 TCP 连接的建立、复用和闲置超时
	transport := &http.Transport{
		// [MaxIdleConns]: 最大闲置连接数
		// 意义：保持跟服务器的长连接，避免频繁三次握手，降低 CPU 消耗
		MaxIdleConns: 100,

		// [IdleconnTimeout]: 一个连接如果没人使用，多久后会自动销毁
		IdleConnTimeout: 90 * time.Second,

		// [DialContext]: 自定义底层拨号逻辑(day 32 知识)
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时
			KeepAlive: 30 * time.Second, // TCP KeepAlive 间隔
		}).DialContext,
	}

	// [http.Client]: 高层的“指挥官”
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second, // 整个请求（包括读取 Body ）的绝对截止时间
	}

	// 发起请求练习
	req, _ := http.NewRequest("GET", "https://api-takumi.mihoyo.com/health", nil)

	// 使用带超时的 Context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("通信成功！收到字节数：%d\n", len(body))
}
