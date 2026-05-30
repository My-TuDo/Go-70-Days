# ===== 第一阶段：编译 =====
# golang:1.26-alpine 是轻量级的 Go 编译环境（~300MB）
# alpine 是 Alpine Linux，体积小
FROM golang:1.26-alpine AS builder

# 设置工作目录
WORKDIR /build

# 先复制依赖文件，利用 Docker 缓存层
# 这两行放在源码之前：只要 go.mod/go.sum 没变，就不会重新下载依赖
COPY go.mod ./
# 没有 go.sum 文件也可以，直接跳过
# COPY go.sum ./

# 下载依赖
RUN go mod download

# 复制所有源码
COPY . .

# 编译成静态二进制
# CGO_ENABLED=0：禁用 C 语言依赖，保证二进制能在任何 Linux 上跑
# -ldflags="-s -w"：去掉调试信息，缩小体积
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /build/log-inspector ./cmd/log-inspector

# ===== 第二阶段：运行 =====
# scratch 是一个完全空的镜像（0MB），只有 Linux 内核调用
# 因为 Go 是静态编译，不需要任何系统库就能跑
FROM scratch

# 从 builder 阶段把编译好的二进制复制过来
COPY --from=builder /build/log-inspector /log-inspector

# 容器启动时执行的命令
ENTRYPOINT ["/log-inspector"]
