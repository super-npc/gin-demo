package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	//ginDemo()
	//slice()
	//deferDemo()
	//deferDemoTest()
	//goroutineTest()
	lockTest()
}

func lockTest() {
	var balance int = 1000
	for i := 0; i < 500; i++ {
		go lockTestThread(&balance, 50)
		go lockTestThread(&balance, -50)
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("执行后: %d", balance)
}

var locker sync.Mutex

func lockTestThread(balance *int, n int) {
	locker.Lock() // 多携程使用互斥锁
	defer locker.Unlock()
	*balance = *balance + n
	fmt.Println(*balance)
}

func goroutineTest() {
	fmt.Println("主协程")

	go func() {
		fmt.Println("子协程")
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			time.Sleep(time.Second * 1)
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("主协程退出") // 只要主协程退出,那么子协程不管是否已经执行完成都会退出
}

func deferDemoTest() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	fmt.Println("c")
	return
	defer fmt.Println("d")
}

func deferDemo() {
	// 适合做close处理,例如:关闭文件流
	fmt.Printf("后进先出 \n")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Printf("后进先出 \n")
}

func fileTest(path string) *os.File {
	open, err := os.Open(path)
	defer func(open *os.File) { // 使用堆栈关闭文件流
		err := open.Close()
		if err != nil {
			panic(err)
		}
	}(open)
	if err != nil {
		fmt.Println("open file error:", err)
	}
	return open
}

func slice() {
	var s = make([]int, 0)
	fmt.Printf("%#v 长度:%d,容量:%d \n", s, len(s), cap(s))
	for i := 0; i < 6; i++ {
		s = append(s, i)
	}
	fmt.Printf("%#v 长度:%d,容量:%d \n", s, len(s), cap(s))
	for i := 0; i < 6; i++ {
		s = append(s, i)
	}
	fmt.Printf("%#v 长度:%d,容量:%d \n", s, len(s), cap(s))
}
