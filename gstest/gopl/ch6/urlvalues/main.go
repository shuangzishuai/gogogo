// Nil也是一个合法的接收器类型
// 就像一些函数允许nil指针作为参数一样，方法理论上也可以用nil指针作为其接收器，尤其当nil对于对象来说是合法的零值时，比如map或者slice
package main

import (
	"fmt"
	"net/url"
)

/*

	//包 net/url
	package url

	// Values maps a string key to a list of values.
	type Values map[string][]string
	// Get returns the first value associated with the given key,
	// or "" if there are none.
	func (v Values) Get(key string) string {
		if vs := v[key]; len(vs) > 0 {
			return vs[0]
		}
		return ""
	}
	// Add adds the value to key.
	// It appends to any existing values associated with key.
	func (v Values) Add(key, value string) {
		v[key] = append(v[key], value)
	}


*/

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	m = nil
	fmt.Println(m.Get("item"))
	/*
		因为传入的是存储了内存地址的变量，你改变这个变量本身是影响不了原始的变量的，想想C语言，是差不多的
	*/
	m.Add("item", "3") // panic: assignment to entry in nil map

}
