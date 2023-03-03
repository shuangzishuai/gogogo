package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//程序将一直阻塞，直至收到 worker 使用通道发送的通知。
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
