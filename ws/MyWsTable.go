package ws

import (
	"base/tool"

	"github.com/labstack/gommon/log"
)

var wsTable = tool.Table[string, string, *MyWsCnf]{}

func TestWsMap() {
	var cm = "spot"

	log.Info(putVal(cm, "BTC-USDT"))
	log.Info(putVal(cm, "LTC-USDT"))
}

func putVal(instType string, instId string) *MyWsCnf {
	v, ok := wsTable.Get(instType, instId)
	if !ok {
		// 初始化客户端
		instant := NewInstant("wss://ws.bitget.com/v2/ws/public")
		v = &instant
		wsTable.Put(instType, instId, v)
	}
	return v
}
