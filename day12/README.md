# Day12 — 网络探测与扫描

> **核心目标**：利用 net 包实现并发端口扫描及站点延迟探测。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **端口扫描** | `net.DialTimeout("tcp", addr, 2s)` | 尝试建立 TCP 连接判断端口是否开放 |
| 2 | **Context 拨测** | `d.DialContext(ctx, "tcp", site)` | 带超时控制的网络拨号，防僵尸连接 |
| 3 | **结果封装** | `type Result struct { Site string; Latency time.Duration }` | 标准化探测结果，便于统计排序 |
| 4 | **并发收集** | `resultChan := make(chan Result, len(sites))` | 使用带缓冲通道收集协程结果 |
| 5 | **延迟排序** | `sort.Slice(results, func(i, j int) bool { ... })` | 对探测结果按耗时升序排列 |

---

## 🧪 运行方式

```bash
cd day12 && go run day12_multi_monitor.go
```
