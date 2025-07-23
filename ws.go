package main

import (
	"bytes"
	"image"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
)

func connectWebSocket() {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"} // 修改为你的 WebSocket 地址
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		_, data, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			log.Println("decode:", err)
			continue
		}

		eImg := ebiten.NewImageFromImage(img)

		queueMutex.Lock()
		if len(imgQueue) > 100 {
			imgQueue = imgQueue[1:]
		}
		imgQueue = append(imgQueue, eImg)
		queueMutex.Unlock()
	}
}
