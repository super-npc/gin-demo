package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
	"ws/bitget"

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
	wsUrl       string
	instType    string
	symbols     []Symbol
	Proxy       Proxy
	MyWsRunning *MyWsRunning
}

type Proxy struct {
	Enable bool
	Url    string
}

type MyWsRunning struct {
	Running     bool
	Conn        *websocket.Conn
	LastMsg     chan string // 最后的数据
	LastRecTime time.Time
}

// NewInstant 2. 唯一构造函数，强制传 wsUrl
func NewInstant(wsUrl string, instType string, symbols []Symbol) MyWsCnf {
	return MyWsCnf{wsUrl: wsUrl, instType: instType, symbols: symbols}
}

func (cnf *MyWsCnf) Start() {
	cnf.connection()
	// 连接成功,发送订阅
	cnf.subscribe()
	// 开启阻塞监听
	go cnf.listen()
}

func (cnf *MyWsCnf) subscribe() {
	// todo 不同交易所不同实现,需要作为多肽
	// 构造消息体
	var reqBitGet = bitget.ReqBitGet{} // 发送消息
	reqBitGet.Op = bitget.Subscribe
	subChannel := "books15"
	var subSymbol = make([]bitget.SubBitGet, 0)
	for _, symbol := range cnf.symbols {
		sub := bitget.SubBitGet{}
		sub.Channel = subChannel
		sub.InstType = cnf.instType
		sub.InstId = fmt.Sprintf("%s%s", symbol.coin, symbol.quote)
		subSymbol = append(subSymbol, sub)
	}
	reqBitGet.Args = subSymbol

	marshal, _ := json.Marshal(reqBitGet)
	jsonReq := string(marshal)
	log.Printf("订阅bitget: %s", jsonReq)

	running := cnf.MyWsRunning
	err := running.Conn.WriteMessage(websocket.TextMessage, []byte(jsonReq)) // "Hello, WebSocket!"
	if err != nil {
		// 需要重连, 发送失败而已,重连不需要在这里做
		log.Println("write error:", err)
		return
	}
}

func (cnf *MyWsCnf) listen() {
	// todo 监听数据,要做不同交易所的实现,每个交易所收到的数据不同
	for {
		running := cnf.MyWsRunning
		_, message, err := running.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			//break
			continue
		}
		fmt.Println("Received:", string(message))
	}
}

func (cnf *MyWsCnf) connection() {
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}
	proxy := cnf.Proxy
	if proxy.Enable {
		proxyURL, _ := url.Parse(proxy.Url) // "http://127.0.0.1:7890"
		dialer.Proxy = http.ProxyURL(proxyURL)
	}
	conn, _, err := dialer.Dial(cnf.wsUrl, nil) // ws://localhost:8080/ws
	if err != nil {
		log.Fatal("dial error:", err)
		return
	}
	// 下面已经夯住监听,不会执行到关闭
	//defer func(conn *websocket.Conn) {
	//	err := conn.Close()
	//	if err != nil {
	//		log.Fatal("关闭异常", err)
	//	}
	//}(conn)

	// 初始化运行时结构
	cnf.MyWsRunning = &MyWsRunning{Running: true, Conn: conn}
}
