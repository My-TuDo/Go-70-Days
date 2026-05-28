# Day20 — IO 流与信号处理

> **核心目标**：掌握 `io.Copy` 高效文件克隆及 `signal.Notify` 捕获系统退出信号。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **IO 流拷贝** | `io.Copy(destFile, source)` | 从 Reader 到 Writer 的数据搬运，底层自动分片 |
| 2 | **信号监听** | `signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)` | 订阅操作系统信号，如 Ctrl+C / kill |
| 3 | **通道阻塞** | `sig := <-sigChan` | 主程序阻塞等待信号，实现优雅关闭 |
| 4 | **资源清理** | 收到信号后执行保存/关闭操作再退出 | 确保数据不丢失，资源正确释放 |

---

## 🧪 运行方式

```bash
cd day20 && go run day20_graceful.go
# 按 Ctrl+C 观察优雅退出流程
```
