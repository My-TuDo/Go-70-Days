# TODO: 多阶段构建 Go 应用，暴露 50051 端口

# 构建阶段
FROM golang:1.26-alpine AS builder

# 设置工作目录
WORKDIR /build

# 复制 go.mod 和 go.sum 文件
COPY go.mod ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构筑应用
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /build/nxproxy ./cmd/api/

# 运行阶段 ： 使用更小的基础镜像
## alpine 5MB左右，有 Shell、包管理器
## busybox 1MB左右，有 Shell
## scratch 0MB，空镜像， 适合GO二进制文件
FROM scratch

# 从构建阶段复制二进制文件
COPY --from=builder /build/nxproxy /nxproxy

# 暴露 gRPC 端口
EXPOSE 50051

# 设置入口点
ENTRYPOINT ["/nxproxy"]

