package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// 1.[Notifier]：抽象化告警
// 作用：通用的告警接口，支持多种通知方式（邮件、短信、Webhook等）
type Notifier interface {
	Send(msg string) error
}

// 2.[WebhookNotifier]：实现具体的通知方式
type WebhookNotifier struct {
	URL string
}

func (w *WebhookNotifier) Send(msg string) error {
	// 构造 JSON
	payload := map[string]string{"text": msg}
	data, _ := json.Marshal(payload)

	// 模拟向飞书发送 POST 请求
	resp, err := http.Post(w.URL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// 3.[自愈函数]：模拟故障自动处理
func autoRemedy(host string) {
	fmt.Printf("[Self-Healing] 检测到节点 %s 异常， 正在尝试远程清理残留并重启...\n", host)
	// 模拟远程执行命令（实际可使用 SSH 或 API 调用）
}

func main() {
	// 假设这三从 Health Checker 传回来的结果
	badResults := []string{"http://127.0.0.1:8080/health", "http://localhost:8080/status"}

	// 初始化警告器
	// 用 httpbin.org 模拟接收 Webhook
	alertHub := &WebhookNotifier{URL: "https://httpbin.org/post"}

	fmt.Println("正在处理异常节点...")

	for _, host := range badResults {
		// 1.发送告警
		msg := fmt.Sprintf("节点故障：%s 无法访问，请及时处理!", host)
		err := alertHub.Send(msg)
		if err != nil {
			fmt.Printf("告警发送失败: %v\n", err)
		} else {
			fmt.Printf("告警已通过 Webhook 推送至运维群：%s\n", host)
		}

		// 2.执行自愈
		autoRemedy(host)
	}

	fmt.Println("执行完毕。")
}
