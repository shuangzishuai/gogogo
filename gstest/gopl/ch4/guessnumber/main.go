package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	fmt.Println("我想了一个1到100之间的数字，你猜是多少？")

	var guess int
	for i := 1; i <= 10; i++ {
		fmt.Printf("第%d次猜，请输入你的答案：", i)
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("你猜的数字太小了")
		} else if guess > target {
			fmt.Println("你猜的数字太大了")
		} else {
			fmt.Printf("恭喜你，你在第%d次猜中了！\n", i)
			return
		}
	}

	fmt.Println("很遗憾，你没有在规定次数内猜中。")
	fmt.Printf("正确答案是%d。\n", target)
}
