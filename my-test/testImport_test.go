package my_test

import (
	one_test "my-test/one-test"
	"testing"
	"time"
)

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
