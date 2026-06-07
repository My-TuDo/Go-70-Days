# NetDiag — 网络诊断工具箱

一个 CLI 工具，对目标地址做一整套网络检查，覆盖 DNS 解析、TCP 端口、HTTP 状态、TLS 证书，输出诊断报告。

---

## 项目结构

```
netdiag/
├── cmd/netdiag/main.go          # 入口：参数解析 + goroutine 并发检查
├── internal/
│   ├── checker/
│   │   ├── result.go            # 检查结果类型定义
│   │   ├── dns.go               # DNS 解析检查
│   │   ├── tcp.go               # TCP 端口连通性检查
│   │   ├── http.go              # HTTP GET 探测
│   │   └── tls.go               # TLS 证书检查
│   └── reporter/
│       └── reporter.go          # 终端输出 + JSON 序列化
├── Makefile
└── README.md
```

---

## 快速开始

```bash
# 编译
make build

# 运行诊断
go run ./cmd/netdiag baidu.com
```

输出示例：

```
=== 网络诊断报告 ===
✅ DNS 解析 | baidu.com | 111.63.65.247 (12ms)
✅ TCP 检查 | baidu.com:443 | 端口开放 (16ms)
✅ TLS 证书 | baidu.com | 2026-08-10 过期 (145ms)
✅ HTTP 检查 | https://baidu.com | HTTP 200 (367ms)
──────────────────────────────────────────────────
```

---

## 检查项

| 检查 | 函数 | 说明 |
|:-----|:-----|:------|
| DNS 解析 | `CheckDNS` | 域名→IP 解析，`net.LookupHost` |
| TCP 端口 | `CheckTCP` | 端口是否开放，`net.DialTimeout` |
| HTTP 探测 | `CheckHTTP` | HTTP 状态码，`http.Client` |
| TLS 证书 | `CheckTLS` | 证书过期时间，`crypto/tls` |

---

## 并发模型

四项检查通过 goroutine + channel 并行执行，总耗时取决于最慢的一项。

```go
results := make(chan checker.Result, 4)

go func() { results <- checker.CheckDNS(host) }()
go func() { results <- checker.CheckTCP(host, "443") }()
go func() { results <- checker.CheckHTTP("https://" + host) }()
go func() { results <- checker.CheckTLS(host) }()
```
