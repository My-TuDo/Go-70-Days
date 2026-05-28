# Day27 — TCP Echo Server

> **核心目标**：手动实现 TCP Echo Server，支持高并发连接处理。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **监听端口** | `net.Listen("tcp", "0.0.0.0:9000")` | 在指定地址端口上开启 TCP 监听 |
| 2 | **接受连接** | `listener.Accept()` | 阻塞等待客户端连接，返回 conn 对象 |
| 3 | **并发处理** | `go handleConnection(conn)` | 每个连接分配独立协程，互不阻塞 |
| 4 | **逐行读取** | `bufio.NewReader(conn).ReadString('\n')` | 使用缓冲读取器按行读取数据 |
| 5 | **Echo 回复** | `conn.Write([]byte(reply))` | 向客户端写入响应数据 |

---

## 🧪 运行方式

```bash
cd day27 && go run .
# 另开终端: nc 127.0.0.1 9000
```
