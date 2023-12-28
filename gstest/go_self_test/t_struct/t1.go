package main

import (
	"fmt"
	"unsafe"
)

type Set[K comparable] map[K]struct{}

func (s Set[K]) Add(val K) {
	s[val] = struct{}{}
}

func (s Set[K]) Remove(val K) {
	delete(s, val)
}

func (s Set[K]) Contains(val K) bool {
	_, ok := s[val]
	return ok
}

func main() {
	//占用空间
	var a int
	var b string
	var c struct{} //空结构体 占用空间为0

	fmt.Println(a, b, c)                                              //0  {}
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c)) //8 16 0

	//地址
	var e struct{}  //空结构体的地址相同
	var e2 struct{} //空结构体的地址相同

	fmt.Printf("%p\n", &e)  //0x1008614c0
	fmt.Printf("%p\n", &e2) //0x1008614c0

	//无状态
	//由于空结构体不包含任何字段，因此它不能有状态。这使得空结构体在表示无状态的对象或情况时非常有用。

	//使用场景
	// - 实现Set集合类型
	// - 用于通道信号
	// - 作为方法接收器

	//实现Set集合类型
	useSet()

	//用于通道信号

}

func useSet() {
	set := Set[string]{}
	set.Add("shuai")
	fmt.Println(set.Contains("shuai")) //true

	set.Remove("shuai")
	fmt.Println(set.Contains("shuai")) //false
}
