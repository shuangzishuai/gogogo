package main

import "fmt"

// 原地修改slice1
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// 原地修改slice2
func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"a", "", "b"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}
