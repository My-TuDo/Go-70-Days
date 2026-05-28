# Day40 — ErrGroup 并发错误管理

> **核心目标**：使用 `errgroup` 管理并发任务，一人报错全组撤退。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **创建组** | `group, ctx := errgroup.WithContext(context.Background())` | 创建一个带取消 Context 的并发组 |
| 2 | **启动任务** | `group.Go(func() error { return checkServer(ctx, id) })` | 替代 `go func()`，自动管理 Add/Done |
| 3 | **错误传播** | `group.Wait()` 返回第一个错误 | 任一任务返回 error 后自动取消其他任务 |
| 4 | **连锁取消** | 错误触发 → `ctx.Done()` → 其他协程收到撤退信号 | 一人失败，全员退出 |

---

## 🧪 运行方式

```bash
cd day40/errgroup && go run .
```
