package channel

type MyChannel1Dto struct {
	Name        string
	InputDataCh chan *InputReq
}

func NewRobotManager(name string) *MyChannel1Dto {
	r := &MyChannel1Dto{
		Name:        name,
		InputDataCh: make(chan *InputReq, 2),
	}
	return r
}

type InputReq struct {
	Req  InputParam
	Resp chan Resp
}

type InputParam struct {
	OrderID string
}

type Resp struct {
	//Code int    `json:"code"`
	Msg string `json:"msg"`
	//Data interface{} `json:"data"`
}
