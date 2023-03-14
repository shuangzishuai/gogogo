package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 不关心格式，只想看值
func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], "==="))
	//fmt.Println(os.Args[0])
	//
	//for i, v := range os.Args {
	//	fmt.Println(i, v)
	//}
	end := time.Now()
	fmt.Println(end.Sub(start)) //26.6µs

}
