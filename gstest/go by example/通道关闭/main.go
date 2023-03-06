package main

import (
	"fmt"
)

/**
关闭 一个通道意味着不能再向这个通道发送值了。
*/

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		j, more := <-jobs
		if more {
			fmt.Println("received job", j)
		} else {
			fmt.Println("received all jobs")
			done <- true
			return
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job")
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
