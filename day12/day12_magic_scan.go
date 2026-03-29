package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scanPort(protocol, hostname string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := net.JoinHostPort(hostname, fmt.Sprintf("%d", port))
	conn, err := net.DialTimeout(protocol, address, 2*time.Second)

	if err != nil {
		// 端口不通的话通常不需要打印
		return
	}
	conn.Close()
	fmt.Printf("发现开放端口“%d\n", port)
}

func main() {
	hostname := "api-takumi.mihoyo.com"
	ports := []int{80, 443, 8080, 8443} // 常见的 HTTP/HTTPS 端口

	var wg sync.WaitGroup
	start := time.Now()

	fmt.Printf("开始并发扫描 %s ...\n", hostname)

	for _, port := range ports {
		wg.Add(1)
		go scanPort("tcp", hostname, port, &wg)
	}

	wg.Wait()
	fmt.Printf("扫描完成，耗时: %v\n", time.Since(start))
}
