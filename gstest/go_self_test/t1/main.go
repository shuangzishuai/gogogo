package main

import (
	"fmt"
	"strings"
)

func main() {
	f := myMin(1.1, 3)
	fmt.Println(f)

	f1 := myAdd(1, 2.2)
	fmt.Println(f1)
}

func tt() {
	str := "123:456"
	str = "123456"
	splitArr := strings.Split(str, ":")
	for i, s := range splitArr {
		fmt.Println(i, s)
	}
}

func myMin[T int | float64](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

func myAdd[T int | float64](a, b T) T {
	return a + b
}
