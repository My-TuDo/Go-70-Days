# Day26 — 外部命令执行

> **核心目标**：使用 `exec.Command` 调用 Linux 系统命令并捕获输出。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **执行命令** | `exec.Command("uptime")` | 定义要执行的外部命令 |
| 2 | **捕获输出** | `cmd.Output()` | 执行命令并捕获标准输出 |
| 3 | **运行检查** | `cmd.Run()` | 执行命令，只关心是否成功（err == nil） |
| 4 | **结果处理** | `strings.TrimSpace(string(output))` | 清理命令输出中的多余空白字符 |

---

## 🧪 运行方式

```bash
cd day26 && go run day26_command.go
```
