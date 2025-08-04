package assets

import (
	"image"
	"image/color"
	_ "image/png"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func LoadImage(path string) *ebiten.Image {
	// 从文件系统读取图片
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open image %s: %v", path, err)
		// 返回一个默认的黑色图片作为占位符
		img := ebiten.NewImage(800, 600)
		img.Fill(color.RGBA{0, 0, 0, 255})
		return img
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Printf("Failed to decode image %s: %v", path, err)
		// 返回一个默认的黑色图片作为占位符
		defaultImg := ebiten.NewImage(800, 600)
		defaultImg.Fill(color.RGBA{0, 0, 0, 255})
		return defaultImg
	}

	ebitenImg := ebiten.NewImageFromImage(img)
	return ebitenImg
}

func LoadFont(path string, size float64) font.Face {
	// 尝试从文件系统读取字体
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open font %s: %v", path, err)
		// 使用默认字体
		tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
		if err != nil {
			log.Printf("Failed to parse default font: %v", err)
			return nil
		}
		face, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    size,
			DPI:     72,
			Hinting: font.HintingFull,
		})
		if err != nil {
			log.Printf("Failed to create font face: %v", err)
			return nil
		}
		return face
	}
	defer file.Close()

	// 读取字体文件内容
	fontBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Failed to read font file %s: %v", path, err)
		// 使用默认字体
		tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
		if err != nil {
			log.Printf("Failed to parse default font: %v", err)
			return nil
		}
		face, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    size,
			DPI:     72,
			Hinting: font.HintingFull,
		})
		if err != nil {
			log.Printf("Failed to create font face: %v", err)
			return nil
		}
		return face
	}

	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Printf("Failed to parse font %s: %v", path, err)
		// 使用默认字体
		tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
		if err != nil {
			log.Printf("Failed to parse default font: %v", err)
			return nil
		}
		face, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    size,
			DPI:     72,
			Hinting: font.HintingFull,
		})
		if err != nil {
			log.Printf("Failed to create font face: %v", err)
			return nil
		}
		return face
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Printf("Failed to create font face: %v", err)
		return nil
	}

	return face
}

func LoadFontWithSize(path string, size float64) font.Face {
	// 尝试从文件系统读取字体
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open font %s: %v", path, err)
		// 使用默认字体
		tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
		if err != nil {
			log.Printf("Failed to parse default font: %v", err)
			return nil
		}
		face, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    size,
			DPI:     72,
			Hinting: font.HintingFull,
		})
		if err != nil {
			log.Printf("Failed to create font face: %v", err)
			return nil
		}
		return face
	}
	defer file.Close()

	// 读取字体文件内容
	fontBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Failed to read font file %s: %v", path, err)
		// 使用默认字体
		tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
		if err != nil {
			log.Printf("Failed to parse default font: %v", err)
			return nil
		}
		face, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    size,
			DPI:     72,
			Hinting: font.HintingFull,
		})
		if err != nil {
			log.Printf("Failed to create font face: %v", err)
			return nil
		}
		return face
	}

	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Printf("Failed to parse font %s: %v", path, err)
		// 使用默认字体
		tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
		if err != nil {
			log.Printf("Failed to parse default font: %v", err)
			return nil
		}
		face, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    size,
			DPI:     72,
			Hinting: font.HintingFull,
		})
		if err != nil {
			log.Printf("Failed to create font face: %v", err)
			return nil
		}
		return face
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Printf("Failed to create font face: %v", err)
		return nil
	}

	return face
}
