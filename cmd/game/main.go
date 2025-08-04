package main

import (
	"ebiten-demo/internal/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Game Main Menu Demo")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
