# Day49 — 文件哈希校验与远程编排

> **核心目标**：远程下发 `sha256sum` 进行文件指纹比对，构造远程命令链完成部署。

---

## 📌 知识点归纳

### hash_verifier（文件哈希校验）

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **SHA256 计算** | `sha256.New()` + `io.Copy(h, f)` | 分块读取计算哈希，不占内存 |
| 2 | **远程校验** | `ssh session.Output("sha256sum /path")` | 远程执行 Linux 原生命令获取哈希 |
| 3 | **指纹比对** | `localHash == remoteHash` | 双重确认文件完整性，不一致则触发清理 |

### remote_orchestrator（远程命令编排）

| # | 概念 | 关键代码 | 一句话说明 |
|---|------|----------|-----------|
| 1 | **命令链** | `[ -f file ] && mkdir -p dir && gunzip -c file > out` | 用 `&&` 串联多个命令，一步执行 |
| 2 | **执行回执** | `session.CombinedOutput(cmd)` | 获取命令执行的全部输出 |
| 3 | **后置验证** | `session.Output("ls | wc -l")` | 验证部署结果，统计解压文件数 |

---

## 🧪 运行方式

```bash
cd day49/hash_verifier && go run .
cd day49/remote_orchestrator && go run .
```
