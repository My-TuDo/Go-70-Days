# Process Reaper — 进程收割者

> **核心目标**：确保子进程及其衍生进程在超时或被中断时被彻底清理，杜绝孤儿进程残留。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **超时控制** | `context.WithTimeout(ctx, 10s)` + `defer cancel()` | 到期自动触发取消，`ctx.Err()` 返回 `DeadlineExceeded` |
| 2 | **Context 驱动进程** | `exec.CommandContext(ctx, ...)` | 超时/取消时自动向子进程发 `SIGKILL`，Go 推荐做法 |
| 3 | **进程组隔离** | `SysProcAttr{Setpgid: true}` | 子进程及其衍生进程划入独立组，**防止逃逸为孤儿** |
| 4 | **信号监听** | `signal.Notify(ch, SIGINT, SIGTERM)` | 捕获 `Ctrl+C` 和终止信号，带缓冲通道防丢信号 |
| 5 | **批量杀进程组** | `syscall.Kill(-PID, SIGKILL)` | Linux 特性：向负 PID 发信号 = 杀掉整个进程组 |
| 6 | **异步启动+同步等待** | `cmd.Start()` → `cmd.Wait()` | `Start` 不阻塞，`Wait` 回收资源防僵尸进程 |
| 7 | **超时 vs 异常判断** | `ctx.Err() == DeadlineExceeded` | 区分超时终止与其他错误，差异化处理 |

---

## 🧪 运行方式

```bash
cd day52/process_reaper && go run main.go
```

- 默认 **10 秒超时**，超时后自动清理子进程
- 按 `Ctrl+C` 可**手动触发**进程组清理并退出

---

## 🔄 执行流程

```mermaid
graph TD
    A[启动 main] --> B[创建 10s 超时 Context]
    A --> C[启动信号监听协程]
    A --> D[创建带 Context 的子进程<br/>设置 Setpgid]
    D --> E[cmd.Start 异步启动]
    E --> F[cmd.Wait 等待结束]
    F --> G{检查 ctx.Err}
    G -->|DeadlineExceeded| H[输出超时提示]
    G -->|其他错误| I[输出异常信息]
    G -->|nil| J[输出正常结束]
    C -->|收到 SIGINT/SIGTERM| K[syscall.Kill -PID<br/>杀掉整个进程组]
    K --> L[os.Exit 退出]
