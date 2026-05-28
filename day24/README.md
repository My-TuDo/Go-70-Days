# Day24 — Init 初始化与 Nil 陷阱

> **核心目标**：掌握 `init` 函数执行顺序及接口 nil 不等于 nil 的经典陷阱。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **init 顺序** | 全局变量 → init() → main() | 包初始化严格遵循此顺序 |
| 2 | **多 init** | 一个文件可写多个 `init()` 函数 | 按代码书写顺序依次执行 |
| 3 | **接口 nil 陷阱** | `var err *MyError = nil; return err` | 返回 nil 指针给 error 接口，接口非 nil（有类型信息） |
| 4 | **类型判断** | `fmt.Printf("%T, %v", err, err)` | 接口为 nil 需同时类型和值都为 nil |

---

## 🧪 运行方式

```bash
cd day24 && go run day24_init_flow.go
```
