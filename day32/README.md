# Day32 — 进程管道与网络拨号

> **核心目标**：掌握 `cmd.StdoutPipe()` 实时截获子进程输出，以及定制化 `net.Dialer`。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **标准输出管道** | `cmd.StdoutPipe()` | 获取子进程的标准输出，实时读取 |
| 2 | **异步启动** | `cmd.Start()` | 非阻塞启动子进程，立即返回 |
| 3 | **同步等待** | `cmd.Wait()` | 等待子进程结束后再继续执行 |
| 4 | **定制拨号器** | `&net.Dialer{Timeout: 5s, KeepAlive: 30s}` | 自定义超时和 KeepAlive 参数的 TCP 拨号器 |
| 5 | **Context 拨号** | `dialer.DialContext(ctx, "tcp", target)` | 带上下文控制的网络连接建立 |

---

## 🧪 运行方式

```bash
cd day32/process_pipe && go run .
cd day32/tcp_dialer && go run .
```
