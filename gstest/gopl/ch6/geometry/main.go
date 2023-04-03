package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

// 普通函数
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Point类型的方法
func (p Point) Distance2(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 路径的长度
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) scaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{6, 4},
		// {1, 1},
	}
	fmt.Println(perim.Distance())

	r := &Point{1, 2}
	r.scaleBy(2)
	fmt.Println(r)

	p := Point{1, 2}
	pptr := &p
	pptr.scaleBy(2)
	fmt.Println(p)

	//如果接收器p是一个Point类型的变量，并且其方法需要一个Point指针作为接收器
	p.scaleBy(3)
	fmt.Println(p)
}
