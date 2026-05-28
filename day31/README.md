# Day31 — 日志扫描与文件锁

> **核心目标**：使用 `bufio.Scanner` 逐行扫描日志文件，`syscall.Flock` 实现进程级单例锁。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **逐行扫描** | `bufio.NewScanner(file).Scan()` | 按行读取大文件，内存占用小 |
| 2 | **关键字过滤** | `strings.Contains(line, "ERROR")` | 扫描过程中筛选特定模式的行 |
| 3 | **文件锁** | `syscall.Flock(fd, LOCK_EX\|LOCK_NB)` | 进程级独占锁，防止任务重复执行 |
| 4 | **非阻塞尝试** | `LOCK_NB` 标志 | 拿不到锁立即返回错误，不等待 |
| 5 | **手动解锁** | `syscall.Flock(fd, LOCK_UN)` | 任务完成后主动释放文件锁 |

---

## 🧪 运行方式

```bash
cd day31/bufio_scanner && go run .
cd day31/day31_file_lock && go run .
```
