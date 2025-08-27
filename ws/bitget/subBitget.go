package bitget

const (
	Subscribe   = "subscribe"
	Unsubscribe = "unsubscribe"
)

type ReqBitGet struct {
	Op   string      `json:"op"`
	Args []SubBitGet `json:"args"`
}

type SubBitGet struct {
	InstType string `json:"instType"`
	Channel  string `json:"channel"`
	InstId   string `json:"instId"` // "instId": "BTCUSDT"
}
