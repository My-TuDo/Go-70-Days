# Day43 — Zap 高性能日志

> **核心目标**：使用 Uber 出品的 `zap` 日志库替代标准库，实现高性能结构化日志。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **生产级日志器** | `zap.NewProduction()` | 创建 JSON 格式、含时间戳/级别/行号的高性能日志器 |
| 2 | **刷新缓冲区** | `defer logger.Sync()` | 程序退出前将内存日志刷入磁盘，防丢失 |
| 3 | **上下文继承** | `logger.With(zap.String("service", "sign-worker"))` | 创建带固定字段的子日志器，每行自动附带 |
| 4 | **强类型字段** | `zap.String()` / `zap.Duration()` / `zap.Int()` | 避免反射开销，比 fmt.Sprintf 快得多 |

---

## 🧪 运行方式

```bash
cd day43/zap_logger && go run .
```
