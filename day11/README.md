# Day11 — HTTP POST 请求

> **核心目标**：掌握使用 `http.Client` 进行 POST 请求及 Header 配置。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **JSON 序列化** | `json.Marshal(data)` | 将 map 转为 JSON 字节流作为请求体 |
| 2 | **构建请求** | `http.NewRequest("POST", url, body)` | 手动构造 HTTP 请求，支持自定义方法/Header |
| 3 | **设置 Header** | `req.Header.Set("Content-Type", "application/json")` | 在请求头中声明内容格式 |
| 4 | **发送请求** | `client.Do(req)` | 执行 HTTP 请求并获取响应 |

---

## 🧪 运行方式

```bash
cd day11 && go run .
```
