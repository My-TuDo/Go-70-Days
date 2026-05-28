# Day51 — 资源深度采集与可视化看板

> **核心目标**：采集系统深度指标，用 ANSI 颜色 + `tabwriter` 渲染终端监控报表。

---

## 📌 知识点归纳

### resource_scraper（资源深度采集）

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **磁盘采集** | `exec.CommandContext(ctx, "df", "/", "--output=pcent")` | 执行 df 命令获取磁盘使用率 |
| 2 | **TCP 采集** | `ss -ant | grep ESTAB | wc -l` | 获取当前 ESTABLISHED 状态的 TCP 连接数 |
| 3 | **正则提取** | `regexp.MustCompile(\`(\d+)\`).FindString()` | 从命令输出中提取数字 |

### visual_dashboard（可视化看板）

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **颜色编码** | `\033[32m` 绿色 / `\033[33m` 黄色 / `\033[31m` 红色 | 通过颜色快速传递健康状态 |
| 2 | **表格对齐** | `tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)` | 自动对齐列宽，生成整洁的终端表格 |
| 3 | **条件染色** | `Disk > 90 → 红色, Disk > 80 → 黄色` | 根据阈值动态选择颜色 |

---

## 🧪 运行方式

```bash
cd day51/resource_scraper && go run .
cd day51/visual_dashboard && go run .
```
