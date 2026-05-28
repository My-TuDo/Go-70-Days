# Day44 — SRE 综合监控引擎

> **核心目标**：集成 Context、Zap、Prometheus 的企业级拨测哨兵。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **拨测引擎** | `dialer.DialContext(ctx, "tcp", target)` | 带超时控制的 TCP 连接探测，防僵尸连接 |
| 2 | **指标标签** | `probeTotal.WithLabelValues(target, "success").Inc()` | 按目标和状态分类统计探测结果 |
| 3 | **日志集成** | `logger.Error("探测失败", zap.String("target", target), zap.Error(err))` | 成功/失败均记录结构化日志 |
| 4 | **Prometheus 网关** | `http.Handle("/metrics", promhttp.Handler())` | 独立端口暴露指标供 Prometheus 抓取 |
| 5 | **巡检循环** | `for { for _, target := range targets { ... }; time.Sleep(10s) }` | 每 10 秒轮询所有目标 |

---

## 🧪 运行方式

```bash
cd day44/sre_engine && go run .
# 访问 http://localhost:8081/metrics
```
