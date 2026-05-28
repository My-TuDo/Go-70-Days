# Day23 — 内存分析与 PProf

> **核心目标**：理解 Go 堆栈分配机制与 pprof 性能分析工具的使用。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **栈上分配** | `a := 10` | 直接定义的局部变量通常在栈上分配 |
| 2 | **堆上分配** | `n := &Node{Val: 100}` | 使用指针且可能逃逸的变量在堆上分配 |
| 3 | **PProf 接入** | `import _ "net/http/pprof"` | 空导入注册 pprof handler |
| 4 | **性能分析** | 访问 `http://localhost:6060/debug/pprof/` | 实时查看 heap / goroutine / cpu 等指标 |

---

## 🧪 运行方式

```bash
cd day23 && go run day23_pprof.go
# 访问 http://localhost:6060/debug/pprof/ 查看性能数据
```
