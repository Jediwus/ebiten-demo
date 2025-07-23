package main

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func mockImageStream() {
	rand.Seed(time.Now().UnixNano())

	for {
		// 模拟生成一张 960x540 的随机彩色图像
		img := image.NewRGBA(image.Rect(0, 0, 960, 540))
		fillRandomColor(img)

		eImg := ebiten.NewImageFromImage(img)

		queueMutex.Lock()
		if len(imgQueue) > 100 {
			imgQueue = imgQueue[1:]
		}
		imgQueue = append(imgQueue, eImg)
		queueMutex.Unlock()

		time.Sleep(time.Millisecond * 33) // 约 30FPS
	}
}

func fillRandomColor(img *image.RGBA) {
	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
}
