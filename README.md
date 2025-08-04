# 🎮 Ebiten 游戏主菜单演示

一个基于 [Ebiten](https://ebiten.org/) 游戏引擎构建的现代化游戏主菜单演示项目，采用 Windows 10 风格设计，支持自适应布局和全屏功能。

## ✨ 功能特性

- 🎯 **现代化界面**: Windows 10 风格的设计语言
- 🖱️ **智能交互**: 支持鼠标点击和悬停效果
- 🎨 **自适应布局**: 界面元素根据窗口大小自动调整
- 🖼️ **智能背景**: Windows 10 风格的壁纸填充效果
- 🔄 **全屏支持**: F11 键切换全屏/窗口模式
- 📱 **响应式设计**: 支持任意窗口尺寸
- ⚡ **高性能**: 60FPS 流畅运行
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
go run ./cmd/game/main.go
```

或者构建可执行文件：

```bash
go build -o game-menu ./cmd/game/main.go
./game-menu
```

## 🎮 使用方法

1. **启动程序**: 运行 `go run ./cmd/game/main.go` 启动游戏菜单
2. **查看效果**: 程序会显示一个现代化的游戏主菜单界面
3. **交互操作**:
   - 点击 "Start Game" 按钮会在控制台输出 "Switch to game scene"
   - 点击 "Quit" 按钮会退出程序
   - 鼠标悬停在按钮上会有视觉反馈
   - 按 F11 键切换全屏/窗口模式
4. **窗口调整**: 可以自由调整窗口大小，界面会自动适应

## 🎨 设计特色

### 自适应布局

- **标题位置**: 位于屏幕上方 1/4 处
- **按钮分布**: Start Game 在 3/5 处，Quit 在 4/5 处
- **字体大小**: 根据窗口宽度自动调整
- **按钮尺寸**: 宽度为窗口的 1/4，高度为窗口的 1/12

### Windows 10 风格

- **背景填充**: 类似 Windows 壁纸的智能填充效果
- **渐变背景**: 无图片时的优雅渐变背景
- **按钮设计**: 现代化的蓝色主题按钮
- **文字阴影**: 增强可读性的阴影效果

### 全屏功能

- **F11 切换**: 一键切换全屏和窗口模式
- **动态适配**: 全屏时界面元素自动调整
- **状态提示**: 左下角显示全屏操作提示

## 📁 项目结构

```
ebiten-demo/
├── cmd/
│   └── game/
│       └── main.go          # 主程序入口
├── internal/
│   ├── assets/
│   │   └── assets.go        # 资源加载模块
│   ├── game/
│   │   └── game.go          # 游戏逻辑
│   └── ui/
│       └── menu.go          # 现代化菜单UI
├── assets/                 # 资源文件
│   ├── fonts/
│   │   └── Roboto-Black.ttf # 字体文件
│   └── images/
│       └── menu_background.png # 背景图片
├── go.mod                   # Go 模块定义
└── README.md               # 项目说明文档
```

## 🛠️ 技术栈

- **[Ebiten v2](https://github.com/hajimehoshi/ebiten)**: 跨平台 2D 游戏引擎
- **[Go](https://golang.org/)**: 编程语言
- **OpenGL**: 图形渲染
- **Windows 10 设计语言**: 现代化 UI 设计

## 🔧 配置说明

### 窗口设置

在 `cmd/game/main.go` 中可以调整初始窗口大小和标题：

```go
ebiten.SetWindowSize(800, 600)
ebiten.SetWindowTitle("Game Main Menu Demo")
```

### 资源文件

- 字体文件位于 `assets/fonts/Roboto-Black.ttf`
- 背景图片位于 `assets/images/menu_background.png`

## 🎨 自定义

### 修改界面布局

在 `internal/ui/menu.go` 中可以调整界面布局：

```go
// 调整标题位置
y := m.windowHeight / 4

// 调整按钮位置
startY := m.windowHeight * 3 / 5
quitY := m.windowHeight * 4 / 5
```

### 修改颜色主题

```go
// 按钮颜色
buttonColor = color.RGBA{60, 100, 180, 255} // 深蓝色
hoverColor = color.RGBA{80, 120, 200, 255}  // 悬停蓝色
```

### 添加新功能

1. 在 `internal/ui/menu.go` 中添加新的按钮
2. 在 `internal/game/game.go` 中处理按钮回调
3. 根据需要添加新的游戏场景

## 📝 开发说明

### 主要组件

- **Game**: 主游戏逻辑，管理菜单状态和窗口
- **Menu**: 现代化 UI 组件，支持自适应布局
- **Assets**: 资源加载模块，支持动态字体大小

### 自适应设计

- **窗口尺寸**: 实时获取窗口大小并调整布局
- **字体缩放**: 根据窗口宽度自动调整字体大小
- **按钮适配**: 按钮大小和位置随窗口变化
- **背景填充**: 智能的背景图片填充算法

### 扩展建议

- 添加更多菜单选项（设置、帮助等）
- 实现场景切换动画
- 添加音效支持
- 实现主题切换功能
- 添加键盘快捷键支持
- 实现多语言支持

## 📄 许可证

本项目采用 MIT 许可证。
