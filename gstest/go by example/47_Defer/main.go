package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating ")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("closing ")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func writeFile(f *os.File) {
	fmt.Println("writing ")
	// data := []string{"11", "22", "33"}
	data := []byte("hello\n")
	fmt.Fprintln(f, string(data))
}

func main() {
	f := createFile("./47_Defer/testfile")
	defer closeFile(f)
	writeFile(f)
}
