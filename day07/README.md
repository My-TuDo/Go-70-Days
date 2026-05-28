# Day07 — Web 与数据库集成

> **核心目标**：集成 Gin 框架与 GORM ORM 库构建基础 RESTful API。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **路由注册** | `r.GET("/add", handler)` | 定义 HTTP 端点与处理函数的映射 |
| 2 | **自动迁移** | `DB.AutoMigrate(&Users{})` | 根据 Go 结构体自动创建/更新数据库表 |
| 3 | **参数提取** | `c.Query("name")` | 从 URL 查询参数中获取数据 |
| 4 | **ORM 创建** | `DB.Create(&newUser)` | 无需手写 SQL，通过对象操作数据库 |
| 5 | **字符串转整型** | `strconv.Atoi(ageStr)` | 将字符串解析为整型，需处理错误 |

---

## 🧪 运行方式

```bash
cd day07 && go run .
# 需要 MySQL 服务运行，配置 go_70_days 数据库
```
