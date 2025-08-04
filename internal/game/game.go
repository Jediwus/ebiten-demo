package game

import (
	"ebiten-demo/internal/assets"
	"ebiten-demo/internal/ui"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type Game struct {
	currentMenu *ui.Menu
	font        font.Face
	bgImage     *ebiten.Image
}

func NewGame() *Game {
	f := assets.LoadFont("assets/fonts/Roboto-Black.ttf", 24)
	bg := assets.LoadImage("assets/images/menu_background.png")
	menu := ui.NewMenu(f, bg)
	g := &Game{currentMenu: menu, font: f, bgImage: bg}
	menu.OnStart = func() {
		// 切换到游戏场景逻辑
		println("Switch to game scene")
	}
	menu.OnQuit = func() {
		os.Exit(0)
	}
	return g
}

func (g *Game) Update() error {
	g.currentMenu.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentMenu.Draw(screen)
}

func (g *Game) Layout(outsideW, outsideH int) (int, int) {
	// 返回当前窗口大小，支持动态调整
	return outsideW, outsideH
}
