package main

import "fmt"

func t1() {
	const sample = "\\xbd\\xb2\\x3d\\xbc\\x20\\xe2\\x8c\\x98"
	fmt.Println(sample)

	//由于 sample 字符串中包含的部分字节不是有效的 ASCII，甚至不是有效的 UTF-8，所以打印的时候会是乱码。
	const sample1 = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" //
	fmt.Println(sample1)                               //��=� ⌘

	//遍历字符串 得到每一个字节
	for i := 0; i < len(sample1); i++ {
		//单独打印的字节和字符串字面量中16进制转义的字节相匹配
		fmt.Printf("%x ", sample1[i]) //bd b2 3d bc 20 e2 8c 98
	}

	fmt.Printf("%x\n", sample1) //bdb23dbc20e28c98

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample1)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample1)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample1)
}

func main() {
	t1()
}
