# Day38 — CI 流水线模拟

> **核心目标**：模拟 CI 过程，依次执行测试、编译并生成报告。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **测试阶段** | `exec.Command("go", "test", "-v", path)` | 执行单元测试，输出实时显示到控制台 |
| 2 | **编译阶段** | `exec.Command("go", "build", "-o", "sentry_app", src)` | 自动化编译为可执行二进制文件 |
| 3 | **Fail Fast** | `if err != nil { return }` | 测试失败立即中断，不继续构建 |
| 4 | **报告生成** | `json.NewEncoder(file).Encode(report)` | 将流水线结果写入 JSON 报告持久化 |

---

## 🧪 运行方式

```bash
cd day38_ci/ci_pipeline && go run .
# 检查生成的 ci_result.json
```
