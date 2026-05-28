# Day45 — Shell 执行器

> **核心目标**：带超时控制的 Shell 脚本批量执行封装。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **超时强杀** | `exec.CommandContext(ctx, "bash", "-c", cmd)` | 超时后自动发送 SIGKILL，防止脚本卡死 |
| 2 | **输出捕获** | `bytes.Buffer{}` 捕获 Stdout / Stderr | 分别捕获标准输出和错误输出 |
| 3 | **退出码获取** | `exitError.ExitCode()` | 获取 Linux 命令的退出状态码 |
| 4 | **结果封装** | `type CommandResult struct { Stdout, Stderr string; ExitCode int }` | 统一结构体返回，便于后续处理 |

---

## 🧪 运行方式

```bash
cd day45/shell_executor && go run .
```
