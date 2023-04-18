package main

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 散列（hash） 经常用于生成二进制文件或者文本块的短标识。
func main() {
	s := "sha256 this string"
	h := sha256.New()
	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Printf("%x\n", s)
	fmt.Println("%x\n", bs)
}
