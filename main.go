package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// 连接 WebSocket
	// go connectWebSocket()

	// 模拟数据源
	// go mockImageStream() // 👈 用随机彩色图像模拟

	go mockJPEGStream() // 👈 用本地JPEG图像模拟

	// 设置窗口大小和标题
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Video Player")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
