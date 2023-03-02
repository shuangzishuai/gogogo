package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	fmt.Printf("co={name: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describe interface {
		describe() string
	}

	//可以使用带有方法的嵌入结构来赋予接口实现到其他结构上。 因为嵌入了 base ，在这里我们看到 container 也实现了 describer 接口。
	var d describe = co
	fmt.Println("describer:", d.describe())
}
