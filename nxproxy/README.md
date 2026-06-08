# nxproxy — Nginx 反向代理 + Go 后端

Nginx 作为反向代理，接收客户端请求，将 `/api/` 路径转发到 Go 后端，`/` 路径直接返回静态文件。双容器 Docker Compose 编排。

---

## 架构

```
客户端 → Nginx (:8080)
              │
              ├── GET /          → 静态文件 (index.html)
              │
              └── GET /api/*     → 反向代理 → Go 后端 (:50051)
                                       ├── /api/hello   → {"message":"Hello, nxproxy!"}
                                       ├── /api/status  → {"status":"ok","time":"..."}
                                       └── /api/panic   → 502（测试用）
```

---

## 项目结构

```
nxproxy/
├── cmd/api/main.go              # Go 后端 API
├── nginx/default.conf           # Nginx 反向代理配置
├── static/index.html            # 静态文件
├── docker-compose.yml           # 双容器编排
├── Dockerfile                   # Go 应用构建
├── Makefile
└── README.md
```

---

## 快速开始

```bash
# 一键构建并启动
make up

# 验证 Nginx 静态文件
curl http://localhost:8080/

# 验证反向代理到 Go 后端
curl http://localhost:8080/api/hello
curl http://localhost:8080/api/status
```

---

## API 说明

所有 API 通过 Nginx 反向代理访问，路径均为 `/api/*`。

| 端点 | 方法 | 说明 |
|:-----|:----|:------|
| `/api/hello` | GET | 返回欢迎消息 |
| `/api/status` | GET | 返回服务状态和当前时间 |
| `/api/panic` | GET | 模拟崩溃，测试 Nginx 502 处理 |

---

## Nginx 配置要点

| 功能 | 说明 |
|:-----|:------|
| 反向代理 | `proxy_pass http://api:50051` |
| 限流 | `limit_req zone=mylimit burst=5` |
| 真实 IP | `proxy_set_header X-Real-IP $remote_addr` |
| 请求追踪 | `proxy_set_header X-Request-Id $request_id` |
| 日志格式 | 自定义格式，包含 `$request_id` 和 `$upstream_response_time` |

---

## 端口说明

| 服务 | 容器端口 | 宿主机端口 | 说明 |
|:-----|:--------:|:----------:|:-----|
| Nginx | 8080 | 8080 | 入口，接收客户端请求 |
| Go API | 50051 | 50051 | 后端服务，仅内部访问 |

---

## Nginx vs 直连 Go

| 场景 | 直连 Go :50051 | Nginx 反向代理 :8080 |
|:-----|:--------------:|:--------------------:|
| 浏览器访问 | 要输端口号 | 标准端口，不用输 |
| 静态文件 | Go 不擅长 | Nginx 直接返回，性能高 |
| 限流 | 代码里实现 | Nginx 配置一行搞定 |
| 日志分析 | 应用日志 | Nginx 访问日志更详细 |
| 负载均衡 | 需要自己写 | Nginx 自带 `upstream` |
