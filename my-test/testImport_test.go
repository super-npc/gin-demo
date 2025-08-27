package my_test

import (
	"echo/obj"
	one_test "my-test/one-test"
	"testing"
	"time"
)

func Test_myObj(t *testing.T) {
	obj.TestObjData()
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
