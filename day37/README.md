# Day37 — JSON 自动化报告

> **核心目标**：将任务结果格式化并序列化为 JSON 报告文件。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **创建文件** | `os.Create("last_job_report.json")` | 在磁盘上创建新文件用于写入 |
| 2 | **JSON 编码器** | `json.NewEncoder(file)` | 创建编码器，将 Go 结构体直接写入文件流 |
| 3 | **格式化输出** | `encoder.SetIndent("", "  ")` | 设置缩进，生成的 JSON 更易读 |
| 4 | **流式编码** | `encoder.Encode(report)` | 将结构体编码为 JSON 并直接写入文件 |

---

## 🧪 运行方式

```bash
cd day37/json_report && go run .
# 检查生成的 last_job_report.json
```
