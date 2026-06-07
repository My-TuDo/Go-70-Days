# Sentry — 报警聚合推送服务

接收客户端推送的报警 JSON，通过聚合去重筛选出需要推送的有效报警。

---

## 项目结构

```
sentry/
├── .github/workflows/ci.yml      # GitHub Actions CI
├── cmd/sentry/main.go             # 入口
├── internal/
│   ├── aggregator/aggregator.go  # 报警聚合去重
│   └── handler/handler.go        # HTTP 处理器
├── Dockerfile
├── Makefile
└── README.md
```

---

## 快速开始

```bash
# 启动服务
make run-sentry

# 发送测试报警
curl -X POST http://localhost:8080/api/alert \
  -H "Content-Type: application/json" \
  -d '{"title":"测试","level":"warn","source":"prober","message":"hello"}'
```

---

## API

### POST /api/alert

接收报警信息。

**请求体：**

```json
{
    "title": "报警标题",
    "level": "critical",
    "source": "来源",
    "message": "详细描述"
}
```

| 字段 | 类型 | 说明 |
|:-----|:----|:-----|
| title | string | 报警标题 |
| level | string | 级别：info / warn / critical |
| source | string | 来源：prober / prometheus |
| message | string | 详细描述 |

**响应：**

```json
{
    "pushed": true,
    "reason": "new_alert"
}
```

| 字段 | 类型 | 说明 |
|:-----|:----|:-----|
| pushed | bool | true 已推送 / false 重复已过滤 |
| reason | string | new_alert / duplicate_within_window |

---

## 环境变量

| 变量 | 默认值 | 说明 |
|:-----|:------|:-----|
| PORT | 8080 | 监听端口 |
| DEDUP_WINDOW | 300 | 去重窗口（秒） |

---

## 去重逻辑

相同 `title` + `source` 的报警，在窗口时间内重复发送会被忽略。
