package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func mockJPEGStream() {
	frameDir := "frames"

	files, err := filepath.Glob(filepath.Join(frameDir, "*.jpg"))
	if err != nil {
		log.Fatal("failed to read JPEG directory:", err)
	}

	sort.Strings(files)

	if len(files) == 0 {
		log.Fatal("no JPEG files found in frames/")
	}

	fmt.Printf("Found %d JPEG frames\n", len(files))

	for {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				log.Println("open error:", err)
				continue
			}

			img, err := jpeg.Decode(f)
			f.Close()
			if err != nil {
				log.Println("decode error:", err)
				continue
			}

			eImg := ebiten.NewImageFromImage(img)

			queueMutex.Lock()
			if len(imgQueue) > 100 {
				imgQueue = imgQueue[1:]
			}
			imgQueue = append(imgQueue, eImg)
			queueMutex.Unlock()

			// ≈ 30 FPS
			// 2025-07-23 Wed. 14:10:57 🥶 测试 30 fps 时发现会有黑屏卡顿
			// time.Sleep(33 * time.Millisecond)

			// ≈ 60 FPS
			time.Sleep(16 * time.Millisecond)
		}
	}
}
