package ui

import (
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
)

type Menu struct {
	ui           *ebitenui.UI
	bgImage      *ebiten.Image
	OnStart      func()
	OnQuit       func()
	windowWidth  int
	windowHeight int
	isFullscreen bool
	startButton  *widget.Button
	quitButton   *widget.Button
}

func NewMenu(fontFace font.Face, bgImg *ebiten.Image) *Menu {
	menu := &Menu{
		bgImage:      bgImg,
		windowWidth:  800,
		windowHeight: 600,
		isFullscreen: false,
	}

	// 创建现代化按钮
	menu.createButtons()

	return menu
}

func (m *Menu) createButtons() {
	// Start Game 按钮
	m.startButton = CreateModernButton("Start Game", func(args *widget.ButtonClickedEventArgs) {
		if m.OnStart != nil {
			m.OnStart()
		}
	})

	// Quit 按钮
	m.quitButton = CreateModernButton("Quit", func(args *widget.ButtonClickedEventArgs) {
		if m.OnQuit != nil {
			m.OnQuit()
		}
	})

	// 创建垂直布局容器
	buttonContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Spacing(20, 20),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
			}),
		),
	)

	// 将按钮添加到容器中
	buttonContainer.AddChild(m.startButton)
	buttonContainer.AddChild(m.quitButton)

	// 创建主容器
	root := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	root.AddChild(buttonContainer)

	m.ui = &ebitenui.UI{Container: root}
}

func (m *Menu) Update() error {
	// 处理F11全屏切换
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		m.isFullscreen = !m.isFullscreen
		if m.isFullscreen {
			ebiten.SetFullscreen(true)
		} else {
			ebiten.SetFullscreen(false)
		}
	}

	// 更新UI
	m.ui.Update()
	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	// 更新窗口尺寸
	m.windowWidth, m.windowHeight = screen.Bounds().Dx(), screen.Bounds().Dy()

	// 绘制Windows 10风格背景
	m.drawBackground(screen)

	// 绘制标题
	m.drawTitle(screen)

	// 绘制UI
	m.ui.Draw(screen)

	// 绘制全屏提示
	m.drawFullscreenHint(screen)
}

func (m *Menu) drawBackground(screen *ebiten.Image) {
	if m.bgImage != nil {
		// Windows 10风格背景填充
		bgWidth, bgHeight := m.bgImage.Bounds().Dx(), m.bgImage.Bounds().Dy()

		// 计算填充模式（类似Windows壁纸的填充效果）
		scaleX := float64(m.windowWidth) / float64(bgWidth)
		scaleY := float64(m.windowHeight) / float64(bgHeight)

		// 使用较大的缩放比例确保覆盖整个屏幕
		scale := scaleX
		if scaleY > scaleX {
			scale = scaleY
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)

		// 居中显示
		scaledWidth := float64(bgWidth) * scale
		scaledHeight := float64(bgHeight) * scale
		offsetX := (float64(m.windowWidth) - scaledWidth) / 2
		offsetY := (float64(m.windowHeight) - scaledHeight) / 2
		op.GeoM.Translate(offsetX, offsetY)

		screen.DrawImage(m.bgImage, op)
	} else {
		// 渐变背景
		m.drawGradientBackground(screen)
	}
}

func (m *Menu) drawGradientBackground(screen *ebiten.Image) {
	// 创建渐变背景
	for y := 0; y < m.windowHeight; y++ {
		ratio := float64(y) / float64(m.windowHeight)
		r := uint8(20 + int(ratio*30))
		g := uint8(30 + int(ratio*40))
		b := uint8(50 + int(ratio*60))
		color := color.RGBA{r, g, b, 255}

		// 绘制水平线
		for x := 0; x < m.windowWidth; x++ {
			screen.Set(x, y, color)
		}
	}
}

func (m *Menu) drawTitle(screen *ebiten.Image) {
	title := "Game Main Menu"

	// 简单的文本宽度计算
	titleWidth := len(title) * 12 // 估算宽度

	x := (m.windowWidth - titleWidth) / 2
	y := m.windowHeight / 4

	// 绘制标题背景
	titleBg := ebiten.NewImage(titleWidth, 30)
	titleBg.Fill(color.RGBA{0, 0, 0, 128})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y-20))
	screen.DrawImage(titleBg, op)
}

func (m *Menu) drawFullscreenHint(screen *ebiten.Image) {
	x := 10
	y := m.windowHeight - 20

	// 全屏时使用更小的位置
	if m.isFullscreen {
		x = 20
		y = m.windowHeight - 30
	}

	// 绘制提示背景
	hintBg := ebiten.NewImage(200, 20)
	hintBg.Fill(color.RGBA{0, 0, 0, 80})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y-15))
	screen.DrawImage(hintBg, op)
}
