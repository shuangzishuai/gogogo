package main

import "fmt"

func main() {
	sum := 0
	nums := []int{2, 3, 4}
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for index, num := range nums {
		if num == 3 {
			fmt.Println("index = ", index)
		}
	}
	kvs := make(map[string]string)
	kvs = map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s --> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println(k)
	}

	//range 在字符串中迭代 unicode 码点(code point)。 第一个返回值是字符的起始字节位置，然后第二个是字符本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
