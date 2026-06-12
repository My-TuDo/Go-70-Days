# 编译阶段，使用官方的 Golang 镜像作为基础镜像
FROM golang:1.26-alpine AS builder

# 设置工作目录
WORKDIR /build

# 先复制 go.mod 和 go.sum 文件，利用 Docker 的缓存机制
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源码
COPY . .

# 编译成二进制
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /build/kvault ./cmd/server/

# 运行阶段，使用一个极小的基础镜像
FROM scratch

# 从 builder 阶段复制编译好的二进制文件
COPY --from=builder /build/kvault /kvault

# 设置容器启动时执行的命令
ENTRYPOINT ["/kvault"]