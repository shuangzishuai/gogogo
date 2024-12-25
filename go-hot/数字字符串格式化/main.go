package main

import (
	"fmt"
	"strings"
)

func solution(s string) string {
	// PLEASE DO NOT MODIFY THE FUNCTION SIGNATURE
	// write code here

	s = strings.Trim(s, "0")

	parts := strings.Split(s, ".")
	integerPart := parts[0]
	var decimalPart string
	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	}

	// 3. 处理整数部分，插入千分位逗号
	var result strings.Builder
	for i, char := range integerPart {
		if i > 0 && (len(integerPart)-i)%3 == 0 {
			result.WriteString(",")
		}
		result.WriteRune(char)
	}

	// 4. 拼接结果
	result.WriteString(decimalPart)

	return result.String()
}

func main() {
	/**
	小M在工作时遇到了一个问题，他需要将用户输入的不带千分位逗号的数字字符串转换为带千分位逗号的格式，并且保留小数部分。小M还发现，有时候输入的数字字符串前面会有无用的 0，这些也需要精简掉。请你帮助小M编写程序，完成这个任务。
	*/
	fmt.Println(solution("1294512.12412") == "1,294,512.12412")
	fmt.Println(solution("0000123456789.99") == "123,456,789.99")
	fmt.Println(solution("987654321") == "987,654,321")
}
