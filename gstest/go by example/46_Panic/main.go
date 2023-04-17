package main

import "os"

func main() {
	panic("a problem") //当函数返回我们不知道如何处理的错误时，终止操作

	_, err := os.Create("./file")

	if err != nil {
		panic(err)
	}
}
