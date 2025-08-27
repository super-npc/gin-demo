package my_test

import (
	one_test "my-test/one-test"
	"testing"
	"time"
	"ws"
)

func Test_ws(t *testing.T) {
	ws.TestWsMap()
}

func Test_myObj(t *testing.T) {

}

func Test_AAA(t *testing.T) {
	myPrint()
}

func Test_orm(t *testing.T) {
	myOrm()
}

func Test_channel(t *testing.T) {
	one_test.MyChannelDemo()
	time.Sleep(time.Second * 10)
}

func Test_table(t *testing.T) {
	collTable()
}
