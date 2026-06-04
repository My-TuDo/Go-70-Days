# Go-70-Days


# 📚 我的后端与 SRE 学习日志

这个仓库用于记录我的日常代码练习与技术积累。所有具体的技术笔记、代码实现和运行指南，均保存在对应的子文件夹中。

---

## 📂 学习目录导航

请点击下方链接查看每天的详细笔记与代码：

### 🧱 基础语法与数据结构
- [📅 Day 01 - 变量与基础数据类型](./day01)
- [📅 Day 02 - 切片应用与循环遍历](./day02)
- [📅 Day 03 - 结构体、指针与方法](./day03)
- [📅 Day 04 - 接口基础与多态实现](./day04)
- [📅 Day 05 - 并发协程与管道挖矿](./day05)
- [📅 Day 06 - JSON 序列化与数据处理](./day06)

### 🌐 Web 开发与网络编程
- [📅 Day 07 - Gin 框架与 HTTP 服务](./day07)
- [📅 Day 11 - HTTP POST 数据提交](./day11)
- [📅 Day 12 - 多站点拨测延迟排序](./day12)
- [📅 Day 27 - TCP Echo 服务端开发](./day27)
- [📅 Day 33 - HTTP 高级客户端定制化](./day33)
- [📅 Day 39 - Gin 中间件与耗时记录](./day39)
- [📅 Day 53 - TCP 连接加固与粘包拆解](./day53)

### ⚙️ 并发编程与系统监控
- [📅 Day 08 - 系统 CPU/内存健康监控](./day08)
- [📅 Day 10 - Context 并发超时控制](./day10)
- [📅 Day 13 - Panic 捕获与系统容错](./day13)
- [📅 Day 17 - Worker Pool 工作池池化](./day17)
- [📅 Day 18 - Select 多路复用监控](./day18)
- [📅 Day 20 - 系统信号监听优雅退出](./day20)
- [📅 Day 22 - Mutex 互斥锁并发安全](./day22)
- [📅 Day 29 - Ticker 心跳定时监测](./day29)
- [📅 Day 30 - Runtime 运行时状态监控](./day30)
- [📅 Day 40 - ErrGroup 并发错误管理](./day40)
- [📅 Day 41 - 令牌桶算法频率限制](./day41)

### 🔧 工程实践与工具链
- [📅 Day 14 - 接口类型断言处理](./day14)
- [📅 Day 15 - 接口高级告警系统实现](./day15)
- [📅 Day 19 - 单元测试基础与实践](./day19_test)
- [📅 Day 21 - 反射机制元编程应用](./day21)
- [📅 Day 23 - PProf 性能瓶颈分析](./day23)
- [📅 Day 24 - Init 注入与初始化流](./day24)
- [📅 Day 25 - 错误包装与上下文追踪](./day25)
- [📅 Day 26 - 系统外部命令调用执行](./day26)
- [📅 Day 28 - Viper 环境变量配置管理](./day28)
- [📅 Day 31 - 逐行日志扫描与文件锁](./day31)
- [📅 Day 32 - 进程管道通信与网络拨号](./day32)
- [📅 Day 37 - 生成 JSON 自动化报告](./day37)

### 🔐 远程运维与自动化
- [📅 Day 34 - SSH 远程单机命令执行](./day34)
- [📅 Day 35 - SSH 并发多机任务下发](./day35)
- [📅 Day 36 - 通道限流与并发控制](./day36)
- [📅 Day 38 - CI 流水线流程模拟](./day38_ci)
- [📅 Day 42 - Prometheus 埋点指标采集](./day42)
- [📅 Day 43 - Zap 高性能日志记录](./day43)
- [📅 Day 44 - SRE 综合监控指标引擎](./day44)
- [📅 Day 45 - 超时控制与 Shell 执行器](./day45)
- [📅 Day 46 - 递归文件目录清理](./day46)
- [📅 Day 47 - Tar/Gzip 压缩包管理](./day47)
- [📅 Day 48 - SFTP 远程制品部署](./day48)
- [📅 Day 49 - 文件哈希校验与远程编排](./day49)
- [📅 Day 50 - 告警自愈与集群巡检](./day50)
- [📅 Day 51 - 资源深度采集与看板展示](./day51)
- [📅 Day 52 - 进程收割与超时回收](./day52)

### 🛠️ 实战项目
- [🔍 log-inspector — 日志巡检 CLI 工具](./log-inspector)
- [📡 prober — 分布式服务拨测系统（高并发 + MySQL + Redis + Docker Compose）](./prober)
- [🔐 kvault — gRPC 配置存储服务（Protobuf + 并发安全 + Makefile）](./kvault)
- [🔐 kvault — gRPC 配置存储服务（Protobuf + gRPC，开发中）](./kvault)

---

## 🚀 统一运行方式

通常情况下，请进入具体目录后运行对应的核心文件：
```bash
cd dayXX
go run main.go  # 或者对应的运行命令