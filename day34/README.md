# Day34 — SSH 远程命令执行

> **核心目标**：通过 Go 远程连接服务器并执行系统命令。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **SSH 客户端配置** | `&ssh.ClientConfig{User, Auth, HostKeyCallback, Timeout}` | 配置连接凭证、超时和主机密钥验证策略 |
| 2 | **建立连接** | `ssh.Dial("tcp", addr, config)` | 通过 TCP 建立 SSH 连接 |
| 3 | **创建会话** | `client.NewSession()` | 在已建立的连接上创建一个新的命令会话 |
| 4 | **执行命令** | `session.CombinedOutput("uptime")` | 执行远程命令并捕获输出 |
| 5 | **资源泄漏演示** | 循环 Dial 不 Close 会耗尽文件描述符 | 演示资源未释放导致的连接泄漏 |

---

## 🧪 运行方式

```bash
cd day34/ssh_cmd && go run .
cd day34/leak_test && go run .
```
