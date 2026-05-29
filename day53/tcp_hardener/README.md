# TCP Hardener — TCP 连接加固

> **核心目标**：利用底层 Socket 控制，加固 TCP 长连接，防御僵尸连接和慢连接攻击。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **底层监听** | `net.ResolveTCPAddr` + `net.ListenTCP` | 获得 `*net.TCPConn` 类型，解锁 Socket 级别的精细控制 |
| 2 | **接收连接** | `listener.AcceptTCP()` | 阻塞等待客户端连接并返回 `*net.TCPConn` |
| 3 | **KeepAlive 保活** | `conn.SetKeepAlive(true)` + `conn.SetKeepAlivePeriod(30s)` | 底层定期发送心跳包，检测异常断电/断网导致的僵尸连接 |
| 4 | **Read Deadline** | `conn.SetReadDeadline(time.Now().Add(10s))` | 每次读取前刷新超时，防御"慢连接攻击"，防止死连接霸占文件描述符 |
| 5 | **超时判断** | `if netErr, ok := err.(net.Error); ok && netErr.Timeout()` | 区分超时断开与客户端主动断开，便于日志区分 |
| 6 | **租约续期** | 每次收到消息 → 回复确认 → 重置 10s Deadline | 活跃连接持续获得有效期延长，沉默自动剔除 |

---

## 🧪 运行方式

```bash
cd day53/tcp_hardener && go run .

# 另开终端测试
nc 127.0.0.1 9000
# 输入任意文字观察"租约续期"效果
# 保持 10 秒不输入观察超时断开
```

---

## 🔄 执行流程

```mermaid
graph TD
    A[启动 TCP 网关 :9000] --> B[AcceptTCP 等待连接]
    B --> C[新玩家连接]
    C --> D[开启 goroutine 处理]
    D --> E[开启 KeepAlive 30s 心跳]
    E --> F[SetReadDeadline 10s 倒计时]
    F --> G{客户端 10s 内发来数据?}
    G -->|是| H[读取数据并回复确认]
    H --> I[重置 ReadDeadline 10s]
    I --> F
    G -->|否 超时| J[netErr.Timeout 判定]
    J --> K[日志记录超时 强制断开]
    K --> L[回收资源 关闭连接]
    G -->|正常断开| M[日志记录正常退出]
    M --> L
```
