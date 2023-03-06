package main

import (
	"fmt"
	"sync"
	"unicode/utf8"
)

/**
使⽤两个goroutine交替打印序列，⼀个goroutine打印数字， 另外⼀个goroutine打印字⺟， 最终效果如下：
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

func printNumLetter() {
	nums, letters := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			<-nums
			fmt.Printf("%d%d", i, i+1)
			i += 2
			letters <- true
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			<-letters
			if i >= utf8.RuneCountInString(str) {
				wait.Done()
				return
			}
			fmt.Print(str[i : i+2])
			i += 2
			nums <- true
		}
	}(&wait)
	nums <- true
	wait.Wait()
}

func main() {
	printNumLetter()
}
