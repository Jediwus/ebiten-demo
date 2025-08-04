package ui

import (
	"bytes"
	"image/color"
	"math"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/gofont/goregular"
)

// DefaultFont 创建默认字体
func DefaultFont() text.Face {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		panic(err)
	}
	return &text.GoTextFace{
		Source: s,
		Size:   20,
	}
}

// drawRoundedRect 在 dst 上绘制一个圆角矩形
func drawRoundedRect(dst *ebiten.Image, x, y, w, h, r float32, clr color.Color) {
	// 中心竖条
	vector.DrawFilledRect(dst,
		x, y+r,
		w, h-2*r,
		clr, true,
	)
	// 中心横条（上部分）
	vector.DrawFilledRect(dst,
		x+r, y,
		w-2*r, r,
		clr, true,
	)
	// 中心横条（下部分）
	vector.DrawFilledRect(dst,
		x+r, y+h-r,
		w-2*r, r,
		clr, true,
	)
	// 四个圆角
	vector.DrawFilledCircle(dst, x+r, y+r, r, clr, true)     // 左上
	vector.DrawFilledCircle(dst, x+w-r, y+r, r, clr, true)   // 右上
	vector.DrawFilledCircle(dst, x+w-r, y+h-r, r, clr, true) // 右下
	vector.DrawFilledCircle(dst, x+r, y+h-r, r, clr, true)   // 左下
}

// DefaultNineSlice 画出未按下状态的九宫格贴图
func DefaultNineSlice(base color.Color) *image.NineSlice {
	const size = 64.0
	const tiles = 16.0
	const radius = 8.0

	tile := float32(size / tiles)
	facet := Mix(base, colornames.Gainsboro, 0.2)

	img := ebiten.NewImage(size, size)

	// 第一层：底色，从 y=tile 到 y=size-tile
	drawRoundedRect(img,
		0, tile,
		size, size-tile,
		radius, base,
	)

	// 第二层：浅色覆盖层，从 y=tile 到 y=size-2*tile
	drawRoundedRect(img,
		0, tile,
		size, size-2*tile,
		radius, facet,
	)

	// 第三层：内层底色，从 x=tile, y=2*tile, 大小=size-2*tile × size-4*tile
	drawRoundedRect(img,
		tile, 2*tile,
		size-2*tile, size-4*tile,
		radius, base,
	)

	// 最后生成可拉伸边界
	return image.NewNineSliceBorder(img, int(tile*4))
}

// PressedNineSlice 画出按下状态的九宫格贴图
func PressedNineSlice(base color.Color) *image.NineSlice {
	const size = 64.0
	const tiles = 16.0
	const radius = 8.0

	tile := float32(size / tiles)
	facet := Mix(base, colornames.Gainsboro, 0.2)

	img := ebiten.NewImage(size, size)

	// 第一层：浅色底，从 0,0 到 size,size
	drawRoundedRect(img,
		0, 0,
		size, size,
		radius, facet,
	)

	// 第二层：内层底色，从 x=tile, y=tile, 大小=size-2*tile
	drawRoundedRect(img,
		tile, tile,
		size-2*tile, size-2*tile,
		radius, base,
	)

	return image.NewNineSliceBorder(img, int(tile*4))
}

// Mix 混合两种颜色
func Mix(a, b color.Color, percent float64) color.Color {
	rgba := func(c color.Color) (r, g, b, a uint8) {
		r16, g16, b16, a16 := c.RGBA()
		return uint8(r16 >> 8), uint8(g16 >> 8), uint8(b16 >> 8), uint8(a16 >> 8)
	}
	lerp := func(x, y uint8) uint8 {
		return uint8(math.Round(float64(x) + percent*(float64(y)-float64(x))))
	}
	r1, g1, b1, a1 := rgba(a)
	r2, g2, b2, a2 := rgba(b)

	return color.RGBA{
		R: lerp(r1, r2),
		G: lerp(g1, g2),
		B: lerp(b1, b2),
		A: lerp(a1, a2),
	}
}

// CreateModernButton 创建现代化按钮
func CreateModernButton(text string, clickedHandler func(*widget.ButtonClickedEventArgs)) *widget.Button {
	return widget.NewButton(
		widget.ButtonOpts.TextLabel(text),
		widget.ButtonOpts.TextFace(DefaultFont()),
		widget.ButtonOpts.TextColor(&widget.ButtonTextColor{
			Idle:    colornames.White,
			Hover:   colornames.White,
			Pressed: Mix(colornames.White, colornames.Black, 0.2),
		}),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:         DefaultNineSlice(colornames.Darkslategray),
			Hover:        DefaultNineSlice(Mix(colornames.Darkslategray, colornames.Mediumseagreen, 0.4)),
			Disabled:     DefaultNineSlice(Mix(colornames.Darkslategray, colornames.Gainsboro, 0.8)),
			Pressed:      PressedNineSlice(Mix(colornames.Darkslategray, colornames.Black, 0.4)),
			PressedHover: PressedNineSlice(Mix(colornames.Darkslategray, colornames.Black, 0.4)),
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				StretchHorizontal:  false, // 不拉伸
			}),
			widget.WidgetOpts.MinSize(200, 48), // 固定宽度200，高度48
		),
		widget.ButtonOpts.ClickedHandler(clickedHandler),
	)
}
