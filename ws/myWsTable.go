package ws

import (
	"base/tool"
	"echo/obj"
	"time"
)

var wsTable = tool.NewTable[string, string, *MyWsCnf]()

func TestWsMap() {
	var cm = "spot"

	initWs(cm, []Symbol{{coin: "btc", quote: "usdt"}})
	initWs(cm, []Symbol{{coin: "ltc", quote: "usdt"}})
	time.Sleep(10 * time.Second)
}

func initWs(instType string, symbols []Symbol) *MyWsCnf {
	v, ok := wsTable.Get(instType, "")
	if !ok {
		// 初始化客户端
		instant := NewInstant("wss://ws.bitget.com/v2/ws/public", obj.UM, symbols)
		instant.Proxy = Proxy{Enable: true, Url: "http://127.0.0.1:7890"}
		v = &instant
		v.Start()
		wsTable.Put(instType, "", v)
	}
	return v
}
