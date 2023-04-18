package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	//Join 还会删除多余的分隔符和目录，使得路径更加规范。
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
