# Day33 — HTTP 高级客户端与重试

> **核心目标**：自定义 `http.Transport` 打造高可靠性客户端，实现指数退避重试。

---

## 📌 知识点归纳

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **连接池** | `MaxIdleConns: 100, IdleConnTimeout: 90s` | 复用长连接，减少三次握手开销 |
| 2 | **KeepAlive** | `DialContext: (&net.Dialer{KeepAlive: 30s}).DialContext` | 维持 TCP 连接活性，防防火墙断开 |
| 3 | **指数退避** | `math.Pow(2, float64(i)) * time.Second` | 等待时间逐次翻倍，避免重试风暴 |
| 4 | **最大重试** | 达到上限后报严重告警 | 防止无限重试消耗资源 |

---

## 🧪 运行方式

```bash
cd day33/http_pro && go run .
cd day33/retry && go run .
```
