package main

import "fmt"

/**
使用recover函数可以从panic中回恢复
recover了可以终止panic终止程序，并让他继续执行
*/

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() { //必须在 defer 函数中调用 recover。
		if r := recover(); r != nil {
			fmt.Printf("Recovered. Error: %v\n", r)
		}
	}()
	mayPanic()
	fmt.Println("after mayPanic()")
}
