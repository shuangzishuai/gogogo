package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println(i)
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("this monday")
	case time.Saturday, time.Sunday:
		fmt.Println("this weekend")
	}

	//不带表达式的 switch 是实现 if/else 逻辑的另一种方式。 这里还展示了 case 表达式也可以不使用常量。
	switch {
	case time.Now().Hour() < 12:
		fmt.Println("没到12点")
	case time.Now().Hour() >= 12:
		fmt.Println("过了12点")
	}

	whoami := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("this is bool")
		case int:
			fmt.Println("this is int")
		case string:
			fmt.Println("this is string")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whoami(15)       //this is int
	whoami(true)     //this is bool
	whoami("hahaha") //this is string
	whoami(17.7)     //Don't know type float64
}
