# Day47 — Tar/Gzip 归档打包

> **核心目标**：利用 `archive/tar` 和 `compress/gzip` 实现文件夹批量备份。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **Gzip 压缩** | `gzip.NewWriter(fw)` | 数据写入时自动压缩，输出到文件 |
| 2 | **Tar 打包** | `tar.NewWriter(gw)` | 记录文件名/权限/目录结构，输出到 Gzip 流 |
| 3 | **文件头写入** | `tw.WriteHeader(header)` | 写入文件的元信息（名称、大小、权限） |
| 4 | **管道链接** | `Tar → Gzip → 磁盘文件` | 数据逐层处理，不占用额外内存 |
| 5 | **遍历打包** | `filepath.Walk` + `io.Copy(tw, fr)` | 遍历目录并将每个文件内容写入 Tar 流 |

---

## 🧪 运行方式

```bash
cd day47/archive_boss && go run .
# 检查生成的 backup_logs.tar.gz
```
