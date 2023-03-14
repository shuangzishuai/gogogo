package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Println(os.Args[0]) ///var/folders/sx/bwcctmzj72lfq887ymkstr5w0000gn/T/go-build1317647436/b001/exe/main
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
