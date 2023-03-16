package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))

	var v1 bool
	var v2 string
	var v3 int
	fmt.Printf("v1 = %v, v2 = %v, v3 = %v", v1, v2, v3)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
