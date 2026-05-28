# Day39 — Gin 中间件与优雅退出

> **核心目标**：Gin 耗时统计中间件 + `http.Server` 优雅关机。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **自定义中间件** | `func StatCost() gin.HandlerFunc { return func(c *gin.Context) { ... } }` | 在请求处理前后注入统一逻辑 |
| 2 | **请求指纹** | `c.Set("trace_id", "trace-123456")` | 在中间件中注入上下文，后续路由可读取 |
| 3 | **耗时监控** | `cost := time.Since(start); if cost > 500ms { 告警 }` | 记录请求耗时，慢接口自动告警 |
| 4 | **手动构建 Server** | `srv := &http.Server{Addr, Handler}` | 替代 `r.Run()`，获取 Server 控制权 |
| 5 | **优雅关闭** | `signal.Notify(quit, SIGINT, SIGTERM); <-quit; srv.Shutdown()` | 收到退出信号后完成正在处理的请求再关闭 |

---

## 🧪 运行方式

```bash
cd day39/gin_middleware && go run .
cd day39/graceful_exit && go run .
```
