package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// 模拟一个很笨的、疯狂申请内存的函数
func memoryLeaker() {
	list := [][]byte{}
	for {
		// 每次申请 1MB 内存并存起来，不释放
		chunk := make([]byte, 1*1024*1024)
		list = append(list, chunk)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// 1. 在后台开启内存泄露协程
	go memoryLeaker()
	fmt.Println("性能分析服务器已启动，访问 http://localhost:6060/debug/pprof/ 来查看性能数据...")
	fmt.Println("正在模拟内存压力，请稍等...")

	// 2. 启动一个简单的 HTTP 服务器来承载监控数据
	// pprof 默认会注册到 http.DefaultServeMux
	http.ListenAndServe(":6060", nil)
}
