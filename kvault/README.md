# 🔐 kvault — gRPC 配置存储服务

一个基于 **gRPC + Protobuf** 的键值对配置存储服务。作为 Go-70-Days 的实战项目之一，旨在练习 gRPC 服务端/客户端开发、接口设计、并发安全、测试与工程化。

---

## 项目状态

✅ **已完成** — gRPC 服务端 + 命令行客户端全功能可用，单元测试 100% 覆盖。

---

## 技术栈

| 层 | 技术 |
|:---|:-----|
| 传输协议 | gRPC (HTTP/2) |
| IDL | Protocol Buffers (proto3) |
| 数据存储 | 内存 KV（`sync.RWMutex` 并发安全） |
| 语言 | Go 1.26 |

---

## 项目结构

```
kvault/
├── api/
│   ├── kvault.proto          # protobuf 接口定义
│   └── kvaultpb/             # 编译生成的 Go 代码
│       ├── kvault.pb.go      # 消息类型（SetRequest, GetResponse 等）
│       └── kvault_grpc.pb.go # 服务端/客户端接口定义
├── cmd/
│   ├── server/main.go        # gRPC 服务端入口
│   └── client/main.go        # CLI 客户端入口
├── internal/
│   ├── server/               # gRPC handler 实现（Controller 层）
│   │   └── server.go         # KVaultServer — 实现 Set/Get/Delete/ListKeys
│   └── store/                # 数据存储层（Model 层）
│       └── store.go          # Store — 内存 KV，RWMutex 并发安全
├── go.mod
└── README.md
```

### 请求链路

```
Client                    Server
  │                         │
  ├── gRPC Set(key,val) ──→├── KVaultServer.Set()
  │                         │   └── store.Set(key, val)
  │                         │       └── mu.Lock() → data[key] = val
  │←──── SetResponse ──────┤
```

---

## API 接口

定义于 [`api/kvault.proto`](./api/kvault.proto)：

| RPC | 请求 | 响应 | 说明 |
|:----|:-----|:-----|:-----|
| `Set` | `key, value` | `success, message` | 写入键值对 |
| `Get` | `key` | `found, value` | 读取键的值 |
| `Delete` | `key` | `success, message` | 删除键值对 |
| `ListKeys` | (空) | `keys[]` | 列出所有键 |

---

## 快速开始

```bash
# 1. 启动 gRPC Server（默认 :50051）
go run ./cmd/server/

# 2. 另一个终端，运行 Client 操作
# 写入
go run ./cmd/client/ set mykey myvalue
# 读取
go run ./cmd/client/ get mykey
# 读取不存在的键
go run ./cmd/client/ get nonexistent
# 删除
go run ./cmd/client/ delete mykey
# 列出所有键
go run ./cmd/client/ list
```

---

## 测试

```bash
# 跑全部测试（含并发安全检测）
go test -v -race ./...

# 查看覆盖率
go test -cover ./internal/store/
go test -cover ./internal/server/
```

当前覆盖率：**store 100%** / **server 100%**

| 包 | 用例数 | 覆盖 |
|:---|:------:|:----:|
| store | 5（含表格驱动 4 条 + 并发安全） | 100% |
| server | 4（Set/Get/Delete/ListKeys + 边界） | 100% |

---

## 构建

```bash
go build ./cmd/server/   # 产出 ./server
go build ./cmd/client/   # 产出 ./client
```

---

## 学习要点

这个项目是为了练习以下知识点：

- ✅ **gRPC 接口实现** — 实现 `kvaultpb.KVaultServer` 接口，理解 `UnimplementedXxxServer` 的作用
- ✅ **gRPC 服务端生命周期** — `net.Listen` → `grpc.NewServer` → `Register` → `Serve`
- ✅ **gRPC 客户端调用** — `grpc.NewClient` → `NewKVaultClient` → 调用 RPC 方法
- ✅ **Response 字段正确填充** — proto 生成的 struct 每个字段都要显式赋值（`Found`、`Success` 等）
- ✅ **优雅退出** — `signal.Notify` + `grpcServer.GracefulStop()`
- ✅ **命令行子命令** — `os.Args` 解析 set/get/delete/list
- ✅ **单元测试 / 集成测试** — store 和 server 双包 100% 覆盖，表格驱动 + 并发安全
- ✅ **Makefile / 工程化收尾** — proto / build / run-server / run-client / clean

---

## 参考

- [Day 10 — Context 并发超时控制](../day10)
- [Day 17 — Worker Pool 工作池池化](../day17)
- [Day 20 — 系统信号监听优雅退出](../day20)
- [Day 07 — Gin 框架与 HTTP 服务](../day07)（prober 的 HTTP 思维可作为 gRPC 对照）
