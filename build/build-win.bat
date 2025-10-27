@echo off
setlocal enabledelayedexpansion

:: ------------------------
:: Go Windows Build Script
:: ------------------------

:: 设置 Windows 编译环境变量
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64

:: 切换到项目目录（基于当前脚本路径）
cd /d "%~dp0\..\server"

:: 清理依赖
echo Running go mod tidy...
go mod tidy

:: 编译 Windows 可执行文件
echo Building for %GOOS%/%GOARCH% ...
go build -o ../build/server.exe -ldflags "-w -s" -trimpath .
if %errorlevel% neq 0 (
    echo Build failed!
    exit /b %errorlevel%
)

:: 压缩（如果 upx 存在）
where upx >nul 2>&1
if %errorlevel% == 0 (
    echo Compressing binary with UPX...
    upx ../build/server.exe
    if %errorlevel% neq 0 (
        echo UPX compression failed, skipping.
    )
) else (
    echo UPX not found, skipping compression.
)

echo Windows executable compiled successfully: ../build/server.exe