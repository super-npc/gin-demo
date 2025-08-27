package ws

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

/*
全局配置
*/
const (
	pingPeriod       = 30 * time.Second // 心跳间隔
	pongWait         = 10 * time.Second // 等待 Pong 的超时
	reconnectInitial = 1 * time.Second  // 第一次重连间隔
	reconnectMax     = 30 * time.Second // 最大重连间隔
)

type MyWsCnf struct {
	wsUrl string

	Proxy struct {
		Enable  bool
		HttpUrl string
	}
	MyWsRunning *MyWsRunning
}

type MyWsRunning struct {
	Running bool
	Conn    *websocket.Conn
}

// NewInstant 2. 唯一构造函数，强制传 wsUrl
func NewInstant(wsUrl string) MyWsCnf {
	return MyWsCnf{wsUrl: wsUrl}
}

func (cnf *MyWsCnf) send(content string) {
	// 发送消息
	running := cnf.MyWsRunning
	err := running.Conn.WriteMessage(websocket.TextMessage, []byte(content)) // "Hello, WebSocket!"
	if err != nil {
		// todo 需要重连
		log.Println("write error:", err)
		return
	}
}

func (cnf *MyWsCnf) listen() {
	for {
		running := cnf.MyWsRunning
		_, message, err := running.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		fmt.Println("Received:", string(message))
	}
}

func (cnf *MyWsCnf) MyWsConn() {
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}
	proxy := cnf.Proxy
	if proxy.Enable {
		proxyURL, _ := url.Parse(proxy.HttpUrl) // "http://127.0.0.1:7890"
		dialer.Proxy = http.ProxyURL(proxyURL)
	}
	conn, _, err := dialer.Dial(cnf.wsUrl, nil) // ws://localhost:8080/ws
	if err != nil {
		log.Fatal("dial error:", err)
	}
	//defer func(conn *websocket.Conn) {
	//	err := conn.Close()
	//	if err != nil {
	//		log.Fatal("关闭异常", err)
	//	}
	//}(conn)

	// 初始化运行时结构
	cnf.MyWsRunning = &MyWsRunning{Running: true, Conn: conn}
	cnf.listen() // 开启阻塞监听
}
