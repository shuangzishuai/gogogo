package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	s := basename("c:\\foo\\bar")
	ss := basename("c:/foo/bar")
	fmt.Println(s)
	fmt.Println(ss)
}
