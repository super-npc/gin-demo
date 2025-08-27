package obj

import (
	"base/tool"

	"github.com/labstack/gommon/log"
)

var tbl = tool.NewTable[string, string, *MyObj]() // 1. 写入

func TestObjData() {
	var cm = CM
	setVal(cm, "BTC-USDT", "test1")
	setVal(cm, "LTC-USDT", "test2")

	log.Info(getVal(cm, "BTC-USDT"))
	log.Info(getVal(cm, "LTC-USDT"))

}

func getVal(instType string, instId string) string {
	v, ok := tbl.Get(instType, instId)
	if !ok {
		return ""
	}
	return v.Name
}

func setVal(instType string, instId string, val string) {
	v, ok := tbl.Get(instType, instId)
	if ok == false {
		v = &MyObj{}
		tbl.Put(instType, instId, v)
	}
	v.Name = val
}
