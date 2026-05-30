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

## 学习笔记

这个项目是第一次**从零构建**一个完整的 Go 项目，要点记录：

- **项目结构**：使用 `cmd/` + `internal/` 分层，分离入口与核心逻辑
- **包设计**：scanner 负责数据采集，reporter 负责输出，职责单一
- **模块化**：类型定义在 `types.go`，逻辑在 `scanner.go`，不混在一起
- **Docker 多阶段构建**：编译阶段用 `golang:alpine`，运行阶段用 `scratch`，缩减最终镜像体积
