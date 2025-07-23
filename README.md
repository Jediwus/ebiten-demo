# 🎮 Ebiten 激光视频播放器演示

一个基于 [Ebiten](https://ebiten.org/) 游戏引擎构建的实时视频播放器演示项目，支持 WebSocket 流媒体和模拟数据源。

## ✨ 功能特性

- 🎯 **实时视频播放**: 基于 Ebiten 2D 游戏引擎的高性能视频渲染
- 🌐 **WebSocket 支持**: 通过网络接收实时视频流数据
- 🎲 **模拟数据源**: 内置随机彩色图像生成器用于测试
- ⚡ **高性能**: 30FPS 流畅播放，支持图像队列缓冲
- 🔧 **易于扩展**: 模块化设计，便于添加新功能

## 🚀 快速开始

### 环境要求

- Go 1.20 或更高版本
- 支持 OpenGL 的图形驱动程序

### 安装依赖

```bash
go mod download
```

### 运行项目

```bash
go run .
```

或者构建可执行文件：

```bash
go build -o laser-player .
./laser-player
```

## 📁 项目结构

```
ebiten-demo/
├── main.go      # 主程序入口
├── plyer.go     # 游戏逻辑和图像渲染
├── ws.go        # WebSocket 连接处理
├── mock.go      # 模拟数据源
├── go.mod       # Go 模块定义
└── README.md    # 项目说明文档
```

## 🔧 配置说明

### WebSocket 连接

在 `ws.go` 文件中修改 WebSocket 服务器地址：

```go
u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
```

### 窗口设置

在 `main.go` 中可以调整窗口大小和标题：

```go
ebiten.SetWindowSize(800, 600)
ebiten.SetWindowTitle("Laser Video Player")
```

## 🎮 使用方法

1. **启动程序**: 运行 `go run .` 启动播放器
2. **查看效果**: 程序会显示一个 800x600 的窗口，播放随机彩色图像
3. **WebSocket 模式**: 取消注释 `main.go` 中的 `connectWebSocket()` 调用以启用网络模式

## 🛠️ 技术栈

- **[Ebiten v2](https://github.com/hajimehoshi/ebiten)**: 跨平台 2D 游戏引擎
- **[Gorilla WebSocket](https://github.com/gorilla/websocket)**: WebSocket 客户端库
- **Go 标准库**: `image`, `sync`, `time` 等

## 🔄 开发模式

### 模拟数据模式（默认）

程序默认运行在模拟数据模式下，会生成随机彩色图像进行播放。

### WebSocket 流模式

要启用 WebSocket 模式，请：

1. 确保有可用的 WebSocket 服务器
2. 修改 `main.go` 中的连接地址
3. 取消注释 `connectWebSocket()` 调用
4. 注释掉 `mockImageStream()` 调用

## 📊 性能特性

- **帧率**: 30 FPS 稳定播放
- **缓冲**: 最多缓存 100 帧图像
- **内存管理**: 自动清理过期帧数据
- **线程安全**: 使用互斥锁保护共享资源

## 🤝 贡献指南

欢迎提交 Issue 和 Pull Request！

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

- [Hajime Hoshi](https://github.com/hajimehoshi) 开发的 Ebiten 引擎
- [Gorilla](https://github.com/gorilla) 提供的 WebSocket 库

---

⭐ 如果这个项目对您有帮助，请给它一个星标！
