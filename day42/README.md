# Day42 — Prometheus 指标采集

> **核心目标**：使用 Prometheus 客户端库定义和暴露监控指标。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **Counter 计数器** | `prometheus.NewCounter(...)` | 累积递增的指标，配合 rate() 看 QPS |
| 2 | **Gauge 压力表** | `prometheus.NewGauge(...)` | 可上下浮动的瞬时值，如在线人数 |
| 3 | **注册指标** | `prometheus.MustRegister(requestCount)` | 将指标注册到默认收集器 |
| 4 | **暴露端点** | `http.Handle("/metrics", promhttp.Handler())` | Prometheus 定期抓取数据 |

---

## 🧪 运行方式

```bash
cd day42/prometheus && go run .
# 访问 http://localhost:8081/metrics 查看指标
```
