#!/usr/bin/env bash
set -e

# ------------------------
# Go Cross-Compile Script
# ------------------------

# 获取脚本所在目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# 设置交叉编译环境
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

# 项目路径
PROJECT_DIR="${SCRIPT_DIR}/../server"
BUILD_DIR="${SCRIPT_DIR}/../build"

echo -e "\033[1;34m[INFO]\033[0m Building Go project for ${GOOS}/${GOARCH}..."
cd "$PROJECT_DIR"

# 确保 build 目录存在
mkdir -p "$BUILD_DIR"

# 下载依赖
echo -e "\033[1;34m[INFO]\033[0m Running go mod tidy..."
go mod tidy

# 编译
echo -e "\033[1;34m[INFO]\033[0m Compiling..."
go build -o "${BUILD_DIR}/server" -ldflags "-w -s" -trimpath .

if [ $? -ne 0 ]; then
    echo -e "\033[1;31m[ERROR]\033[0m Build failed!"
    exit 1
fi

# 检查 UPX
if command -v upx >/dev/null 2>&1; then
    echo -e "\033[1;34m[INFO]\033[0m Compressing binary with UPX..."
    upx "${BUILD_DIR}/server" || echo -e "\033[1;33m[WARN]\033[0m UPX compression failed, skipping."
else
    echo -e "\033[1;33m[WARN]\033[0m UPX not found, skipping compression."
fi

echo -e "\033[1;32m[SUCCESS]\033[0m Build completed successfully!"
echo -e "Output: ${BUILD_DIR}/server"
