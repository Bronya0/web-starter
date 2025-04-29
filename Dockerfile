# ---- 构建阶段 ----
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o /app/server /app/main.go

# ---- 最终阶段 ----
FROM alpine:3.18

# 安装运行时依赖并配置时区
RUN apk update && apk add --no-cache \
    ca-certificates \
    tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && addgroup -S appgroup \
    && adduser -S appuser -G appgroup \
    && rm -rf /var/cache/apk/*

# 创建应用目录并设置权限
WORKDIR /app
RUN chown appuser:appgroup /app

# 复制二进制文件并设置所有权
COPY --from=builder --chown=appuser:appgroup /app/server /app/server

# 切换到非 root 用户
USER appuser

# 暴露应用程序监听的端口
EXPOSE 8080

# 定义运行应用程序的命令。
ENTRYPOINT ["/app/server"]

# 如果你需要传递可以被覆盖的默认参数，请与 ENTRYPOINT 一起使用 CMD：
# CMD ["--default-flag", "value"]