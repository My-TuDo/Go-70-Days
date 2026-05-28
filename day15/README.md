# Day15 — 可插拔告警架构

> **核心目标**：基于接口实现不同告警媒介的随时切换。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **告警接口** | `type Notifier interface { Send(msg string) error }` | 定义通用的通知行为协议 |
| 2 | **控制台实现** | `func (ConsoleNotifier) Send(msg) error` | 实现接口，将消息打印到控制台 |
| 3 | **飞书推送实现** | `func (FeishuNotifier) Send(msg) error` | 模拟通过 Webhook 发送到飞书 |
| 4 | **可插拔调用** | `TriggerAlert(ConsoleNotifier{}, msg)` | 不关心具体实现，只管调用接口方法 |

---

## 🧪 运行方式

```bash
cd day15 && go run .
```
