# Day19 — 单元测试与基准测试

> **核心目标**：掌握 Go 测试框架，编写表格驱动测试与性能测试。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **表格驱动测试** | `cases := []struct{ name string; a,b,expected int }` | 用切片组织多个测试用例，循环执行 |
| 2 | **子测试** | `t.Run(tc.name, func(t *testing.T) { ... })` | 为每个用例创建独立子测试，便于定位失败点 |
| 3 | **错误报告** | `t.Errorf("FAILED: %s", tc.name)` | 记录错误但不中断测试，继续执行后续用例 |
| 4 | **基准测试** | `func BenchmarkAdd(b *testing.B) { for i:=0; i<b.N; i++ }` | 测试函数性能，b.N 由框架自动确定 |

---

## 🧪 运行方式

```bash
# 运行单元测试
go test -v ./day19_test/

# 运行基准测试
go test -bench=. ./day19_test/
```
