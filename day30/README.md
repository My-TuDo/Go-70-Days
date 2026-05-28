# Day30 — Runtime 运行时监控

> **核心目标**：读取 `runtime.MemStats` 获取内存与 GC 累计指标。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **CPU 核心数** | `runtime.NumCPU()` | 获取当前系统的逻辑 CPU 数量 |
| 2 | **协程数量** | `runtime.NumGoroutine()` | 获取当前活跃的 goroutine 总数 |
| 3 | **内存快照** | `runtime.ReadMemStats(&m)` | 读取堆内存分配等运行时统计信息 |
| 4 | **GC 次数** | `m.NumGC` | 查看垃圾回收的累计触发次数 |

---

## 🧪 运行方式

```bash
cd day30 && go run .
```
