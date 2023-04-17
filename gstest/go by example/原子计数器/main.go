package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64 //使用无符号的整型来表示计数器

	var wg sync.WaitGroup //等待所有goroutine结束

	for i := 0; i < 20; i++ { //开启50个goroutine
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
				if atomic.LoadUint64(&ops)%200 == 0 {
					fmt.Println("ops:", ops)
				}
				// ops++ //非原子性操作，最后的结果会得到不同的数字 go run -race main.go
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops)
}
