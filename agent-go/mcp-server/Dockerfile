# 使用官方 Go 基础镜像
FROM golang:alpine AS builder

# 设置工作目录
WORKDIR /build

# 设置环境变量
ENV GO111MODULE=on

# 将当前目录下的所有文件复制到工作目录
COPY . .

# 构建 Go 应用
RUN go build -o mcp-server .

# 使用一个轻量级的 Linux 镜像作为最终镜像
FROM scratch AS deploy

# 设置工作目录
WORKDIR /app

# 将构建好的二进制文件和 vendor 目录从构建阶段复制到最终镜像
COPY --from=builder /build/mcp-server .
COPY --from=builder /build/vendor ./vendor

# 启动应用
CMD ["/app/mcp-server"]