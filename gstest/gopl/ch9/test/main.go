package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("============")
	var x, y int
	go func() {
		x = 1
		fmt.Println("y:", y)
	}()

	go func() {
		y = 1
		fmt.Println("x:", x)
	}()

	time.Sleep(time.Second)

}
