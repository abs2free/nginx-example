# 使用 Go 的官方镜像作为基础镜像
FROM golang:1.23 AS builder

# 设置工作目录
WORKDIR /app

# 首先复制依赖文件
COPY go.mod ./
RUN go mod download

# 复制源代码
COPY . .

# 构建 Go 应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# 使用较小的基础镜像来运行应用程序
FROM alpine:latest

# 添加 ca-certificates 和时区数据
RUN apk --no-cache add ca-certificates tzdata

# 设置工作目录
WORKDIR /root/

# 从 builder 阶段复制构建的应用程序
COPY --from=builder /app/myapp .

# 复制静态文件
COPY --from=builder /app/index.html ./
#COPY --from=builder /app/images ./images/

# 暴露端口
EXPOSE 3000

# 运行应用程序
CMD ["./myapp"]
