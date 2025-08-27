package ws

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func MyWsConn() {
	proxyURL, _ := url.Parse("http://127.0.0.1:7890")
	dialer := websocket.Dialer{
		Proxy:            http.ProxyURL(proxyURL), // 设置 HTTP 代理
		HandshakeTimeout: 10 * time.Second,
	}
	conn, _, err := dialer.Dial("ws://localhost:8080/ws", nil)
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
