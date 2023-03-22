package main

import "fmt"

// 函数向表示十进制非负整数的字符串中插入逗号
//"12345",从右侧开始每三位数宇后面就插人一个逗号,形如 〝12,345”。这个版本仅对整数有效。对浮点数的处理方式留作练习。
//TODO 使用AI提示代码

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Println("test comma")
	s := comma("12345")
	fmt.Println(s)
}
