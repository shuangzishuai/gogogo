package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func f1(v int) *int {
	v = 100
	return &v
}

func main() {
	i, j := 1, 1
	for i <= 3 {
		i++
		var p = f()
		fmt.Printf("f() done p = %v\n", p) //每一次返回的地址是不同的
	}

	v := 1
	for j <= 3 {
		j++
		var p = f1(v)
		fmt.Printf("f1() down p = %v\n", p) //传递变量 每一次返回的地址是不同的
	}
}
