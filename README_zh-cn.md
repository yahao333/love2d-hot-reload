# love2d-hot-reload

一个用 Go（Golang）为 LÖVE2D 项目实现的热重载工具。该工具利用 Go 强大的文件监控和进程管理能力，为 LÖVE2D 项目创建一个高效的开发环境。

## 功能特点

- 实时监控 `.lua` 文件和资源文件的变化
- 自动管理 LÖVE2D 进程
- 跨平台支持（主要在 macOS 上测试）
- 防抖动重载机制，避免频繁重启
- 可配置的文件忽略模式

## 前置要求

- Go 1.22 或更高版本
- 系统中已安装 LÖVE2D
- 基本的 Go 和 LÖVE2D 开发知识

## 安装步骤

1. 克隆仓库：
```bash
git clone https://github.com/yahao333/love2d-hot-reload.git
cd love2d-hot-reload
```

2. 安装依赖：
```bash
go mod download
```

3. 构建项目：
```bash
make
```
或者手动构建：
```bash
go build -o build/hot-reload.bin src/main.go
```

## 使用方法

1. 将 `build/hot-reload.bin` 可执行文件放在你的 LÖVE2D 项目目录中
2. 运行工具：
```bash
./hot-reload.bin
```

工具将会：
- 启动你的 LÖVE2D 项目
- 监控所有相关文件的变化
- 在检测到变化时自动重启 LÖVE2D

## 工作原理

- **文件监控**：使用 `fsnotify` 检测文件系统变化
- **进程管理**：控制 LÖVE2D 进程，在检测到变化时重启
- **防抖动**：实现防抖动机制，防止快速连续重启
- **文件过滤**：忽略临时文件和不相关的变化

## 配置说明

可以通过修改 `src/main.go` 中的以下内容来配置工具：
- 项目目录路径
- 忽略的文件模式
- 防抖动时间（默认：500毫秒）
- LÖVE2D 可执行文件路径

## 贡献指南

欢迎提交 Pull Request 来改进这个项目！

## 许可证

本项目采用 MIT 许可证 - 详见 LICENSE 文件。

## 致谢

- [LÖVE2D](https://love2d.org/) - 本工具为之设计的游戏框架
- [fsnotify](https://github.com/fsnotify/fsnotify) - 文件系统监控库 