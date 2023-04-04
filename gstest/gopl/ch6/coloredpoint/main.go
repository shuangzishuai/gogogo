package main

import (
	"fmt"
	"image/color"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)

	cp.Point.Y = 2
	fmt.Println(cp.Point.Y)

	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	blue := color.RGBA{R: 0, G: 0, B: 255, A: 255}
	p := ColoredPoint{Point: Point{1, 1}, Color: red}
	q := ColoredPoint{Point: Point{2, 2}, Color: blue}
	fmt.Println(p.Point)
	fmt.Println(q.Point)
}
