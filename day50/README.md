# Day50 — 告警自愈与集群巡检

> **核心目标**：并发探测集群 API 状态，实现 Webhook 告警与自愈逻辑。

---

## 📌 知识点归纳

### alert_remedy（告警自愈）

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **告警接口** | `type Notifier interface { Send(msg string) error }` | 抽象化通知方式 |
| 2 | **Webhook 推送** | `http.Post(url, "application/json", body)` | 通过 POST 请求将告警推送到飞书等平台 |
| 3 | **自愈函数** | `autoRemedy(host)` | 检测到异常后自动执行清理和重启 |

### cluster_health（集群健康巡检）

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **Context 超时** | `context.WithTimeout(ctx, 5s)` | 整个巡检任务的总超时控制 |
| 2 | **并发探测** | `go probeHost(ctx, url, resultChan, &wg)` | 同时探测所有节点 |
| 3 | **实时报表** | 格式化打印节点地址、状态码和延迟 | 终端输出结构化的健康总览 |

---

## 🧪 运行方式

```bash
cd day50/alert_remedy && go run .
cd day50/cluster_health && go run .
```
