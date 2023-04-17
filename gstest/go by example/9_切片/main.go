package main

import (
	"fmt"
)

func testSlice() {
	var buffer [256]byte
	fmt.Println(buffer)
	var slice []byte = buffer[100:255]
	//fmt.Println("1 = ", slice)
	//slice = slice[5:10]
	//slice = slice[1 : len(slice)-1]
	//fmt.Println("2 = ", slice)
	//slashPos := bytes.IndexRune(slice, '/')
	//fmt.Println(slashPos)
	//
	//AddOneToEachElement(slice)
	//fmt.Println("3 = ", slice)

	slice = buffer[10:20]
	fmt.Println("4 = ", slice)
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)
}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}

func testLen() {
	var buffer [256]byte
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}

	fmt.Println("Before: len(slice) =", len(slice))
	newSlice := SubtractOneFromLength(slice)
	fmt.Println("After:  len(slice) =", len(slice))
	fmt.Println("After:  len(newSlice) =", len(newSlice))
}

func testLenPtr() {
	var buffer [256]byte
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}

	fmt.Println("Before: len(slice) =", len(slice))
	PtrSubtractOneFromLength(&slice)
	fmt.Println("After:  len(newSlice) =", len(slice))
}

func main() {
	//testSlice()
	//var buffer [256]byte
	//fmt.Println(buffer)
	//var slice []byte = buffer[100:255]
	//fmt.Println("Before: len(slice) =", len(slice))
	//newSlice := SubtractOneFromLength(slice)
	//fmt.Println("After:  len(slice) =", len(slice))
	//fmt.Println("After:  len(newSlice) =", len(newSlice))

	//var buffer [256]byte
	//slice := buffer[10:20]
	//for i := 0; i < len(slice); i++ {
	//	slice[i] = byte(i)
	//}
	//fmt.Println("before", slice)
	//AddOneToEachElement(slice)
	//fmt.Println("after", slice)

	testLen()
	fmt.Println("===============")
	testLenPtr()
}

func t1() {
	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD) //2d: [[0] [1 2] [2 3 4]]
}
