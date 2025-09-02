package waitGroup

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 干完活减 1
	fmt.Printf("worker %d start\n", id)
	time.Sleep(time.Duration(id) * time.Second) // 模拟耗时
	fmt.Printf("worker %d done\n", id)
}

func Test_work(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)         // 增加计数
		go worker(i, &wg) // 注意传指针
	}

	wg.Wait() // 阻塞直到计数器归零
	fmt.Println("all workers finished")
}
