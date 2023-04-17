package main

import (
	"fmt"
	"time"
)

/*
*
我们经常需要在未来的某个时间点运行 Go 代码，或者每隔一定时间重复运行代码。 Go 内置的 定时器 和 打点器 特性让这些变得很简单。 我们会先学习定时器，然后再学习打点器。
*/
func main() {
	//定时器表示在未来某一时刻的独立事件。 你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道。 这里的定时器将等待 2 秒。
	timer1 := time.NewTimer(2 * time.Second)

	//<-timer1.C 会一直阻塞， 直到定时器的通道 C 明确的发送了定时器失效的值。
	<-timer1.C
	fmt.Println("Timer 1 fried")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fried")
	}()
	//time.Sleep(1 * time.Second)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
