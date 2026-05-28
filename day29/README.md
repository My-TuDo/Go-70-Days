# Day29 — Ticker 心跳监测

> **核心目标**：通过 `time.Ticker` 实现定时数据上报。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **定时器** | `time.NewTicker(2 * time.Second)` | 每隔固定时间往 Channel 发送时间戳 |
| 2 | **一次性定时** | `time.After(11 * time.Second)` | 只发送一次信号的 Channel，适合做截止 |
| 3 | **多路复用** | `select { case <-ticker.C: ... case <-done: ... }` | 同时监听定时上报和停止信号 |
| 4 | **资源释放** | `defer ticker.Stop()` | 确保定时器在函数退出时停止，防止泄漏 |

---

## 🧪 运行方式

```bash
cd day29 && go run .
```
