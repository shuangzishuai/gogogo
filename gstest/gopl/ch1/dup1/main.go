package main

import (
	"bufio"
	"fmt"
	"os"
)

// 这个程序用来统计重复的行数
func main() {
	fmt.Println("=====test=====")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
		counts[input.Text()]++
	}
	fmt.Println(counts)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
