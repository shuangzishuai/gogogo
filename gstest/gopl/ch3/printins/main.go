package main

import (
	"bytes"
	"fmt"
	"os"
)

// 将整型slice 转为 字符串输出
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v) //Fprintf根据格式指定符进行格式化并写入w。
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	// The n and err return values from Fprintf are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)

	s := intsToString([]int{1, 2, 3}) //[1, 2, 3]
	fmt.Println(s)
}
