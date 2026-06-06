### Sentry - 报警聚合推送服务
该项目是通过接收客户端传回的 JSON 报文信息，通过匹配和筛选，找到需要进行推送的有效的报警信息

## 项目结构

```
sentry/
├── .github/workflows/ci.yml   # GitHub Actions CI
├── cmd/sentry/main.go          # 入口
├── internal/
│   ├── aggregator/aggregator.go # 报警聚合去重
│   └── handler/handler.go      # HTTP 处理器
├── Dockerfile
├── Makefile
└── README.md
```

## 快速开始

```bash
# 启动服务
make run-sentry

# 发送测试报警
curl -X POST http://localhost:8080/api/alert \
    -H "Content-Type: application/json" \
    -d '{"title":"测试","level":"warn","source":"prober","message":"hello"}'
```


## API

# POST /api/alert
接收报警信息。
请求体：
```json
{
    "title":"报警标题",
    "level":"critical",
    "source":"来源",
    "message":"详细描述"
}
```
## 环境变量

变量          默认值  说明
———————————— —————— ——————————
PORT         8080   监听端口
DEDUP_WINDOW 300    去重窗口(秒)

## 去重逻辑
相同 title + source 的报警，在窗口时间内重复发送会被忽略。