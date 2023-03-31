package main

import "math"

type Point struct {
	X, Y float64
}

// 普通函数
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Point类型的方法
func (p Point) Distance2(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
