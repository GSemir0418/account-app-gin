# 使用官方的 Golang 镜像作为构建阶段的基础镜像
FROM golang:1.21-alpine3.20 AS builder

# 设置工作目录
WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖包
RUN go mod download

# 复制项目文件
COPY . .

# 构建 Gin 应用
RUN go build -o main .

# 使用一个更小的基础镜像来运行应用
FROM alpine:latest

ENV GIN_ENV prod
ENV GIN_MODE release

# 安装必要的包
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /app

# 从构建阶段复制生成的二进制文件
COPY --from=builder /app/main /app/main
COPY --from=builder /app/.env.prod /app/.env.prod 

# 暴露服务端口
EXPOSE 8080

# 运行数据库迁移并启动应用
CMD ["sh", "-c", "./main db migrate:create && ./main server"]