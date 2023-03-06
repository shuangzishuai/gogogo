package main

import (
	"fmt"
	"time"
)

/*
*
超时 对于一个需要连接外部资源， 或者有耗时较长的操作的程序而言是很重要的。 得益于通道和 select，在 Go 中实现超时操作是简洁而优雅的。
*/
func main() {
	c1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "hello"
	}()
	/**
	这里是使用 select 实现一个超时操作。 res := <- c1 等待结果，<-time.After 等待超时（1秒钟）以后发送的值。 由于 select 默认处理第一个已准备好的接收操作， 因此如果操作耗时超过了允许的 1 秒的话，将会执行超时 case。
	*/
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
