# Prober — 分布式服务拨测系统

一个高并发服务健康拨测系统，并发探测多个目标 URL，结果存入 MySQL，最新状态缓存在 Redis，通过 HTTP API 查询。

---

## 🏃 程序执行全流程（建议先看这个）

### 程序启动后，到底发生了什么事？

整个项目由 5 个文件组成，按这个顺序执行：

```
启动命令：go run ./cmd/prober
        │
        ▼
① main.go（入口）
   ├── 读取 config.yaml → 得到 targets 列表
   ├── 连接 MySQL
   ├── 连接 Redis
   ├── 启动 HTTP 服务（调用 handler.SetupRoutes）
   └── 启动后台拨测循环（调用 collector.RunProbeCycle）
   
              ║                     ║
              ║ 每 30 秒            ║ 随时用 curl 访问
              ▼                     ▼
              │                     │
    ② collector.go              ③ handler.go
      RunProbeCycle                 SetupRoutes
        │                            │
        ├── ProbeAll                  ├── GET /health → HealthHandler
        │   └── 并发 5 个 goroutine   ├── GET /status → StatusHandler
        │       └── ProbeOne          └── GET /history → HistoryHandler
        │           └── HTTP GET
        │
        ├── db.Create → 写入 MySQL    handler 反过来读数据：
        │                              │
        └── UpdateCache → 写入 Redis    ├── StatusHandler  → 读 Redis
          (cache.go)                    └── HistoryHandler → 读 MySQL
```

### 关键：数据流向是单向的

```
collector（写）               handler（读）
    │                            │
    ├──→ 写入 MySQL ───────────→  ├── /history 从 MySQL 读
    ├──→ 写入 Redis  ───────────→  ├── /status  从 Redis 读
    │
    拨测是主动的（每 30 秒）        API 是被动的（你访问才响应）
```

### 为什么 SetupRoutes 已经注册了三个路由，但我还没碰过 handler？

因为我在给你的**模板代码里已经写好了**。你看 `handler.go` 里的 `SetupRoutes`：

```go
func SetupRoutes(r *gin.Engine, db *gorm.DB, rdb *redis.Client) {
    r.GET("/health", HealthHandler)          // ← 已经注册好了
    r.GET("/status", func(c *gin.Context) {  // ← 已经注册好了
        StatusHandler(c, rdb)
    })
    r.GET("/history", func(c *gin.Context) { // ← 已经注册好了
        HistoryHandler(c, db)
    })
}
```

这三行就是路由注册。`r.GET("/health", HealthHandler)` 的意思是：**当有人访问 `/health` 时，执行 `HealthHandler` 函数**。

你需要写的是 `HealthHandler`、`StatusHandler`、`HistoryHandler` 这三个函数**内部的具体逻辑**，路由注册已经由模板帮你做好了。

| 代码 | 谁写的 | 作用 |
|:----|:------:|:-----|
| `r.GET("/health", HealthHandler)` | 我（模板） | 告诉 Gin："有人访问 /health 时调 HealthHandler" |
| `func HealthHandler(c *gin.Context) { ... }` | **你填** | 实际返回什么内容 |

### 文件之间的调用关系（谁 import 了谁）

```
cmd/prober/main.go
    ├── import "prober/internal/collector"   → 调 collector.RunProbeCycle
    ├── import "prober/internal/handler"     → 调 handler.SetupRoutes
    └── import "prober/internal/model"       → 用 model.TargetConfig

internal/collector/collector.go
    ├── import "prober/internal/model"       → 用 model.ProbeResult
    └── import "gorm.io/gorm"                → 用 *gorm.DB 写入 MySQL

internal/collector/cache.go
    ├── import "prober/internal/model"       → 用 model.ProbeCache
    └── import "github.com/redis/go-redis/v9" → 用 *redis.Client

internal/handler/handler.go
    ├── import "prober/internal/collector"   → 调 collector.GetCachedStatus
    ├── import "prober/internal/model"       → 用 model.ProbeResult
    ├── import "github.com/redis/go-redis/v9"
    └── import "gorm.io/gorm"
```

**核心规律**：只有 `main.go` 是"总调度"，其他文件只做自己负责的事。`collector` 只负责探测和存数据，`handler` 只负责响应 HTTP 请求，`model` 只定义类型——**谁也别越界**。

---

## 🧠 `c *gin.Context` 是什么？（看完这个你就不蒙了）

每个 handler 函数都有这个参数：

```go
func HealthHandler(c *gin.Context) { ... }
```

### 一句话

`c` 就是**这次 HTTP 请求的"窗口"**——通过它拿请求的数据，也通过它返回响应。

### 拆开看：一次 HTTP 请求的全过程

```
你发请求：curl http://localhost:8080/health
                                    │
                                    ▼
Gin 收到请求，创建一个 gin.Context 对象（就是参数 c）
    │
    ├── c 里装着：
    │     ├── 请求方法 (GET)
    │     ├── URL 路径 (/health)
    │     ├── 查询参数 (?name=baidu)
    │     └── 请求体 (POST 的数据)
    │
    ├── 交给对应的 handler 函数
    │     └── HealthHandler(c)
    │
    └── handler 通过 c 返回响应
          └── c.JSON(200, gin.H{"status": "ok"})
                       │        │
                       │        └── 要返回的数据（Go 的 map）
                       │
                       └── HTTP 状态码（200 = 成功）
```

### 你用过哪些 c 的方法？

| 方法 | 在哪个 handler 里 | 做了什么 |
|:----|:-----------------|:---------|
| `c.JSON(code, data)` | 三个 handler 都用到了 | 返回 JSON 响应 |
| `c.Query("name")` | HistoryHandler | 从 URL 取参数，如 `/history?name=baidu` → 拿到 `"baidu"` |
| `c.Request.URL.Path` | （本项目中没用到） | 拿到请求路径 `/health` |

### `c.JSON` 究竟干了什么？

```go
c.JSON(http.StatusOK, gin.H{"status": "ok"})
```

这行代码做了三件事：

| 步骤 | 实际效果 |
|:----|:---------|
| ① 把 `gin.H{"status": "ok"}` 序列化成 JSON | `{"status":"ok"}` |
| ② 设置 HTTP 响应头 `Content-Type: application/json` | 告诉浏览器"这是 JSON" |
| ③ 设置状态码为 200 | 告诉浏览器"请求成功" |
| ④ 把 JSON 字符串写入响应体 | `{"status":"ok"}` 出现在 curl 的输出里 |

等价于手动写：

```go
c.Writer.Header().Set("Content-Type", "application/json")
c.Writer.WriteHeader(200)
c.Writer.WriteString(`{"status":"ok"}`)
```

所以 `c.JSON` 是一个**快捷方式**，一行顶上面四行。

### `c *gin.Context` 的 `*` 是什么意思？

`*gin.Context` 表示传进来的是**指针**。

```go
func HealthHandler(c *gin.Context) {
    c.JSON(...)      // 这里修改了 c 的内容
}
```

如果不用指针，Gin 传递给 handler 的会是 Context 的**副本**，你在 handler 里设置响应内容，Gin 那边拿不到。

用了指针，**你改的就是 Gin 手里那个**，设了响应 Gin 就能读到。

### 类比

```
HTTP 请求             = 顾客来到前台
gin.Context (c)      = 这个顾客的"接待单"
c.Query("name")      = 从接待单上看顾客填的名字
c.JSON(200, data)    = 在接待单上写好回复，交给前台寄回去
*c                   = 直接改原始接待单（而不是复印件）
```

---

## 项目结构

```
prober/
├── cmd/prober/main.go            # 入口：连接 DB + 启动服务 + 拨测循环
├── internal/
│   ├── model/model.go            # 数据类型定义
│   ├── collector/
│   │   ├── collector.go          # 并发拨测引擎（goroutine pool）
│   │   └── cache.go              # Redis 缓存读写
│   └── handler/handler.go        # HTTP API 处理器
├── config.yaml                   # 探测目标配置
├── go.mod
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

## 快速开始（Docker Compose 一键启动）

### 前置条件

- Docker ≥ 24.0
- Docker Compose v2

### 启动

```bash
cd prober
docker compose up -d
```

此命令会同时启动三个容器：

| 容器 | 作用 | 端口 |
|:----|:----|:----:|
| `prober-mysql-1` | 存储拨测历史记录 | 3306 |
| `prober-redis-1` | 缓存最新状态 | 6379 |
| `prober-prober-1` | 应用本身（拨测 + API） | 8080 |

### 验证

```bash
# 健康检查
curl http://localhost:8080/health

# 查看最新状态（从 Redis 读取）
curl http://localhost:8080/status

# 查看历史记录（从 MySQL 读取）
curl "http://localhost:8080/history?name=baidu"
```

### 停止

```bash
docker compose down
```

### 清理数据

```bash
docker compose down -v   # 加 -v 会删除 MySQL 数据卷
```

---

## 单独启动 MySQL + Redis（不启动应用）

用于开发调试，应用在本机跑，DB 用 Docker：

```bash
docker compose up -d mysql redis
```

然后在本机运行应用：

```bash
go run ./cmd/prober
```

---

## API 说明

### `GET /health`

服务健康检查。

```json
{"status": "ok"}
```

### `GET /status`

返回所有目标的最新拨测状态（从 Redis 读取）。

### `GET /history?name=baidu`

返回指定目标最近 20 条历史记录（从 MySQL 读取）。

---

## 配置说明

编辑 `config.yaml` 添加要探测的目标：

```yaml
targets:
  - name: "baidu"
    url: "https://www.baidu.com"
    timeout_sec: 5
  - name: "github"
    url: "https://github.com"
    timeout_sec: 5
```

---

---

## 🧠 Redis + JSON 缓存知识点（复习用）

### 为什么需要 Redis？直接查 MySQL 不行吗？

可以，但慢。

| 场景 | MySQL | Redis |
|:----|:-----:|:-----:|
| 存历史记录（1 万条） | ✅ 擅长 | ❌ 不适合 |
| 查最新状态（响应 < 1ms） | ❌ 每次查都要走 SQL 解析 | ✅ 内存操作 |
| 频繁读写（每秒 1000 次） | ❌ 磁盘 IO 瓶颈 | ✅ 内存级速度 |

所以本项目**两个都用**：

```
拨测结果
    │
    ├── MySQL：永久存储，供 /history API 查历史
    └── Redis：临时缓存（60 秒过期），供 /status API 秒回
```

### JSON 序列化是什么？

程序里的数据是 Go 结构体（`ProbeCache`），Redis 存的是字符串。中间需要一个"翻译"过程：

```
Go 结构体                       Redis 字符串
┌─────────────────┐           ┌──────────────────────────────────┐
│ ProbeCache {     │  json.   │ {"target_name":"baidu",          │
│   TargetName     │ ────→   │  "success":true,                 │
│   Success: true  │ Marshal  │  "latency_ms":123,...}           │
│   ...            │           │                                  │
└─────────────────┘           └──────────────────────────────────┘

Go 结构体                       Redis 字符串
┌─────────────────┐           ┌──────────────────────────────────┐
│ ProbeCache {     │  json.   │ {"target_name":"baidu",          │
│   TargetName     │ ←────   │  "success":true,                 │
│   ...            │Unmarshal│  ...}                             │
└─────────────────┘           └──────────────────────────────────┘
```

- `json.Marshal`：Go 结构体 → JSON 字符串（写 Redis 时用）
- `json.Unmarshal`：JSON 字符串 → Go 结构体（读 Redis 时用）

### 为什么 JSON 字段有 `json:"target_name"`？

打开 `internal/model/model.go` 你会看到：

```go
type ProbeCache struct {
    TargetName string `json:"target_name"`  // ← 这个叫 struct tag
    TargetURL  string `json:"target_url"`
}
```

`json:"target_name"` 的意思是：**序列化时，字段名用 `target_name` 而不是 `TargetName`。**

```
Go 写法                JSON 写法
TargetName  ───→    "target_name"    (蛇形命名，符合 API 惯例)
TargetURL   ───→    "target_url"
Success     ───→    "success"
```

不加 struct tag 的话，默认会变成大写字母开头的字段名：

```json
{"TargetName":"baidu","TargetURL":"https://..."}  // 不专业
```

所以 struct tag 的作用就是**控制 JSON 里的字段名**。

### `rdb.Set(ctx, key, value, TTL)` 每个参数是什么意思？

```go
rdb.Set(ctx, "probe:status:baidu", jsonData, 60*time.Second)
```

| 参数 | 值 | 含义 |
|:----|:---|:----:|
| `ctx` | `context.Background()` | 告诉 Redis"这是这次请求的上下文"（固定写法） |
| `key` | `"probe:status:baidu"` | 存在 Redis 里的**名字**，类似变量名 |
| `value` | `jsonData` | 存在 Redis 里的**值**，JSON 格式的 `[]byte` |
| `TTL` | `60 * time.Second` | 过期时间，60 秒后 Redis 自动删除这条数据 |

### 什么是 TTL（过期时间）？

TTL = Time To Live（存活时间）。

```
写入 Redis                        60 秒后自动删除
    │                                  │
    ▼                                  ▼
    ┌──────────────────────────────────────┐
    │  probe:status:baidu = {...}          │
    │  还剩 43 秒...                        │
    │  还剩 12 秒...                        │
    │  已过期，自动消失                      │
    └──────────────────────────────────────┘
```

如果没有 TTL，Redis 里的数据会永远占用内存。设了 TTL，**不需要你手动删**，到期自动清理。

### `redis.Nil` 是什么？

当你读一个**不存在的 key** 时，Redis 不会报错，而是返回一个特殊的空值 `redis.Nil`。

```go
jsonData, err := rdb.Get(ctx, key).Bytes()

// 情况 1：key 存在     → err == nil,  jsonData 有数据
// 情况 2：key 不存在    → err == redis.Nil  （不是真正的错误）
// 情况 3：网络断了      → err == timeout/connection refused  （真正的错误）
```

所以在 `GetCachedStatus` 里要分别处理：

```go
if err == redis.Nil {
    return nil, nil    // 没缓存 ≠ 报错，返回 nil 表示"没找到"
}
if err != nil {
    return nil, err    // 真的出错了
}
```

---

## 并发模型

采用 **channel 信号量**模式控制并发数：

```
sem := make(chan struct{}, 5)   // 最多 5 个 goroutine 同时执行

for _, target := range targets {
    sem <- struct{}{}            // 占一个车位（满则阻塞）
    go func(t TargetConfig) {
        defer func() { <-sem }() // 释放车位
        results <- ProbeOne(t)
    }(target)
}
```

参考：`day36/limit_concurrency` 和 `day17/worker_pool`。
