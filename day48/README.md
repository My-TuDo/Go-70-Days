# Day48 — SFTP 远程制品部署

> **核心目标**：通过 SFTP 协议将本地备份包远程推送至部署机。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **SSH 隧道** | `ssh.Dial("tcp", host, config)` | 建立加密的 SSH 连接作为传输通道 |
| 2 | **SFTP 客户端** | `sftp.NewClient(conn)` | 基于 SSH 连接创建文件传输会话 |
| 3 | **流式上传** | `io.Copy(dstFile, srcFile)` | 零拷贝方式将本地文件流写入远程 |
| 4 | **传输统计** | `bytesCopied` 记录传输字节数 | 计算传输速率和耗时 |

---

## 🧪 运行方式

```bash
cd day48/artifact_deployer && go run .
# 需提前在 day47 生成 backup_logs.tar.gz
```
