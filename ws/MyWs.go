package ws

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func MyWsConn() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("关闭异常", err)
		}
	}(conn)

	// 发送消息
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
	if err != nil {
		log.Println("write error:", err)
		return
	}

	// todo 使用协程接收消息
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		fmt.Println("Received:", string(message))
	}
}
