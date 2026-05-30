# log-inspector — 日志巡检工具

一个 Go CLI 工具，递归扫描指定目录下的所有 `.log` 文件，统计各级别日志数量（INFO / WARN / ERROR），并生成 JSON 格式报告。

---

## 项目结构

```
log-inspector/
├── cmd/log-inspector/main.go    # 入口：CLI 参数解析 + 串联流程
├── internal/
│   ├── scanner/
│   │   ├── types.go             # 核心数据类型（LogFileResult, ScanResult, Summary）
│   │   └── scanner.go           # 目录遍历 + 日志文件逐行扫描
│   └── reporter/
│       └── reporter.go          # JSON 报告生成 + 终端摘要打印
├── testdata/                    # 测试用日志文件
│   ├── server.log
│   └── app.log
├── go.mod
├── Dockerfile
└── README.md
```

---

## 快速开始

```bash
# 1. 扫描 testdata 目录
go run ./cmd/log-inspector ./testdata
```

输出示例：

```
=== 日志扫描报告 ===
扫描时间: 2025-07-18 10:30:00
扫描文件: 2 个
总行数:   26 行
-----------------
INFO:  15 行
WARN:   6 行
ERROR:  5 行  ⚠️
报告已保存到: ./testdata_report.json
```

---

## 用法

```bash
go run ./cmd/log-inspector <目标目录>
```

- `<目标目录>`：要扫描的目录路径（必填，不传则默认扫描 `./testdata`）
- 程序会递归遍历该目录下所有后缀为 `.log` 的文件
- JSON 报告会保存在 `<目标目录>_report.json`

---

## Docker 部署

```bash
# 构建镜像
docker build -t log-inspector .

# 运行（挂载日志目录）
docker run --rm -v $(pwd)/testdata:/logs log-inspector /logs
```

---

## 实现功能

| 功能 | 说明 |
|------|------|
| ✅ 递归扫描 | 自动遍历子目录下所有 `.log` 文件 |
| ✅ 三级统计 | 分别统计 INFO / WARN / ERROR 行数 |
| ✅ ERROR 详情 | 记录每个文件前 5 条 ERROR 的具体内容 |
| ✅ JSON 报告 | 生成格式化 JSON 文件，含汇总数据 |
| ✅ 终端摘要 | 运行后直接在终端打印统计概览 |
| ✅ Docker 化 | 多阶段构建，镜像体积 < 10MB |

---

## 🧠 命令行参数知识点（复习用）

### `go run ./cmd/log-inspector ./testdata` 逐词拆解

```
go    run    ./cmd/log-inspector    ./testdata
│     │      │                      │
│     │      │                      └─ 传给 main.go 的参数
│     │      │                         → os.Args[1]
│     │      │
│     │      └─ 包的路径（不是文件名）
│     │         告诉 Go："去这个目录编译并运行"
│     │
│     └─ 子命令：编译 + 运行一步到位
│
└─ Go 工具链入口
```

### `os.Args` 是什么？

`os.Args` 是一个字符串切片，存的是你在终端敲的所有"词"。

```go
// 运行：go run ./cmd/log-inspector ./testdata
//
// os.Args[0] = 编译出的二进制路径（Go 自己管，你看不到）
// os.Args[1] = "./testdata"
// len(os.Args) = 2
```

敲几个词，`os.Args` 就有几个元素。

### 为什么代码里有默认值还要传参？

```go
dir := "./testdata"               // ① 先给一个保底值
if len(os.Args) >= 2 {            // ② 如果你传了参数
    dir = os.Args[1]              // ③ 就用你的，覆盖保底值
}
```

| 你敲的命令 | dir 最终值 | 原因 |
|:-----------|:----------:|:-----|
| `go run ./cmd/log-inspector` | `./testdata` | 没传参，走默认值 |
| `go run ./cmd/log-inspector /var/log` | `/var/log` | 传了参数，覆盖默认值 |

**默认值的作用**：让你不加参数也能跑，而不是报错崩溃。

### 为什么用 `./cmd/log-inspector` 而不是 `main.go`？

这是 Go 的生产级项目结构。一个项目可以有多个入口：

```
cmd/
├── server/      → 启动 HTTP 服务
├── worker/      → 启动后台任务
└── log-inspector/ → 日志巡检工具
```

每个子目录都有自己的 `main.go`，`go run ./cmd/xxx` 选择跑哪一个。
如果所有 `main.go` 都放根目录，Go 不允许同一个包里有多个 `main` 函数。

---

## 学习笔记

这个项目是第一次**从零构建**一个完整的 Go 项目，要点记录：

- **项目结构**：使用 `cmd/` + `internal/` 分层，分离入口与核心逻辑
- **包设计**：scanner 负责数据采集，reporter 负责输出，职责单一
- **模块化**：类型定义在 `types.go`，逻辑在 `scanner.go`，不混在一起
- **Docker 多阶段构建**：编译阶段用 `golang:alpine`，运行阶段用 `scratch`，缩减最终镜像体积
