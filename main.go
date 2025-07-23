package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// 连接 WebSocket
	// go connectWebSocket()

	// 模拟数据源
	go mockImageStream()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Laser Video Player")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
