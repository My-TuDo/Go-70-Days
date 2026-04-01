package main

import (
	"fmt"
)

// 1.定义一个告警接口
type Notifier interface {
	Send(msg string) error
}

// 2.实现 A: 控制台告警
type ConsoleNotifier struct{}

func (n ConsoleNotifier) Send(msg string) error {
	fmt.Printf("[控制台告警] %s\n", msg)
	return nil
}

// 3.实现 B: 飞书机器人（模拟）
type FeishuNotifier struct {
	WebhookURL string
}

func (n FeishuNotifier) Send(msg string) error {
	// 模拟发送消息到飞书机器人
	fmt.Printf("[飞书推送] 发送到 %s: %s\n", n.WebhookURL, msg)
	return nil
}

// 4.核心逻辑： 触发告警（它不关心是谁发，只管调用 Send 方法）
func TriggerAlert(n Notifier, message string) {
	n.Send(message)
}

func main() {
	msg := "米游社 Cookie 即将过期！"

	// 随时切换发送平台，这就是“可插拔” 架构
	TriggerAlert(ConsoleNotifier{}, msg)                                         // 发送到控制台
	TriggerAlert(FeishuNotifier{WebhookURL: "https://feishu.cn/robot/123"}, msg) // 发送到飞书
}
