# ---- 构建阶段 ----
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o /app/server ./cmd/main.go

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

# 从构建阶段复制二进制文件和必要的配置文件
COPY --from=builder --chown=appuser:appgroup /app/server /app/server
COPY --from=builder --chown=appuser:appgroup /app/conf /app/conf

# 切换到非 root 用户
USER appuser

# 暴露应用程序监听的端口
EXPOSE 8080

# 定义运行应用程序的命令。
ENTRYPOINT ["/app/server"]

# 如果你需要传递可以被覆盖的默认参数，请与 ENTRYPOINT 一起使用 CMD：
# CMD ["--default-flag", "value"]