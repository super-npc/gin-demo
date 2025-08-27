package one_test

import (
	"log"
	"time"
)

func MyChannelDemo() {
	ints := make(chan int)

	go func() {
		log.Println("准备数据")
		time.Sleep(time.Second * 2)
		ints <- 10
		log.Println("数据完成")
	}()
	log.Println("准备接收数据")
	log.Println(<-ints)
	log.Println("接收完成")
}
