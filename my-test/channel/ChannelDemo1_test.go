package channel

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_channel(t *testing.T) {
	manager := NewRobotManager("Robot1")

	go func() {
		for {
			aa := <-manager.InputDataCh // 这一步会报错
			fmt.Println(aa)
			//resp := <-aa.Resp
			//fmt.Println(resp)

			i := len(manager.InputDataCh)
			t.Logf("长度: %d", i)
			time.Sleep(time.Second * 2)
		}
	}()

	r := &InputReq{Req: InputParam{OrderID: "111"}, Resp: make(chan Resp, 1)}
	manager.InputDataCh <- r
	//resp1 := <-r.Resp
	fmt.Printf("")

	r2 := &InputReq{Req: InputParam{OrderID: "2222"}, Resp: make(chan Resp, 1)}
	manager.InputDataCh <- r2

	time.Sleep(time.Second * 10)
	//select {}
}

func Test_channel2(t *testing.T) {
	manager := NewRobotManager("Robot1")

	go func() {
		for r := range manager.InputDataCh { // 持续消费
			fmt.Println("请求:", r.Req.OrderID)
			r.Resp <- Resp{Msg: "返回值"}
		}
	}()

	r := &InputReq{Req: InputParam{OrderID: "111"}, Resp: make(chan Resp, 1)}
	manager.InputDataCh <- r
	resp := <-r.Resp
	marshal, _ := json.Marshal(resp)
	fmt.Println("收到处理返回值:", marshal)

	//r2 := &InputReq{Req: InputParam{OrderID: "2222"}, Resp: make(chan Resp, 1)}
	//manager.InputDataCh <- r2

	//time.Sleep(time.Second * 10)
	select {}

}
