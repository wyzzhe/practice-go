# 使用官方 Go 基础镜像
FROM golang:alpine AS builder

# 设置工作目录
WORKDIR ./

# 项目依赖配置文件拷贝到容器中，并下载依赖
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将当前目录下的所有文件复制到工作目录
COPY . .

# 构建 Go 应用
RUN go build -o app .

# 使用一个轻量级的 Linux 镜像作为最终镜像
FROM scratch as deploy

# 设置工作目录
WORKDIR /app

# 将构建好的二进制文件从构建阶段复制到最终镜像
COPY --from=build /build/app .
COPY --from=build /build/conf ./conf

# 启动应用
CMD ["/app"]