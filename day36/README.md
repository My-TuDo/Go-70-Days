# Day36 — 并发控制与结果聚合

> **核心目标**：使用 Channel 和 Semaphore 模式控制并发度并收集任务结果。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **信号量限流** | `sem := make(chan struct{}, maxConcurrency)` | 用带缓冲通道作为令牌桶，控制并发协程数 |
| 2 | **结果封装** | `type ExecResult struct { Host, Output string; Success bool; Cost time.Duration }` | 结构化返回数据，便于主程序统计 |
| 3 | **异步关闭** | `go func() { wg.Wait(); close(resultChan) }()` | 等待所有生产者结束再关闭通道 |
| 4 | **结果消费** | `for res := range resultChan { ... }` | 主协程循环读取通道汇总所有结果 |

---

## 🧪 运行方式

```bash
cd day36/limit_concurrency && go run .
cd day36/result_aggregator && go run .
```
