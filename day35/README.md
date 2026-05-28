# Day35 — SSH 并发多机任务

> **核心目标**：批量对多台服务器并行下发指令。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **批量并发** | `for _, ip := range ips { wg.Add(1); go executeRemoteCommand(ip, ...) }` | 对多个目标同时发起 SSH 连接并执行命令 |
| 2 | **WaitGroup 管理** | `wg.Add(1)` / `wg.Done()` / `wg.Wait()` | 等待所有并发任务完成后统一退出 |
| 3 | **闭包陷阱** | `go func(targetIP string, taskID int) { ... }` | 将循环变量作为参数传递，避免引用同一变量 |

---

## 🧪 运行方式

```bash
cd day35/ssh_concurrent && go run .
```
