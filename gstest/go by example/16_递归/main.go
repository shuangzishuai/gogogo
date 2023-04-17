package main

import "fmt"

func fact(i int) int {
	if i == 0 {
		return 1
	}
	return i * fact(i-1)
}

func main() {
	r := fact(7)
	fmt.Println(r)

	r = fact(5)
	fmt.Println(r)

	var fib func(int2 int) int //闭包函数也可以是递归的,但是要求定义闭包之前用类型化的var 显示声明闭包
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(3))
	fmt.Println(fib(5))
	fmt.Println(fib(7))
}
