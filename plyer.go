package main

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	imgQueue   []*ebiten.Image
	queueMutex sync.Mutex
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	queueMutex.Lock()
	defer queueMutex.Unlock()

	if len(imgQueue) > 0 {
		current := imgQueue[0]
		imgQueue = imgQueue[1:]
		screen.DrawImage(current, nil)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 960, 540
}
