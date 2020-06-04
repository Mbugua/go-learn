package main

import "fmt"

type shape interface {
	area() int
}

type square struct {
	sideLength int
}

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	a := r.height * r.width
	return a
}
func (s square) area() int {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Print(s.area())
}

func main() {
	// areaRectangle := rectangle{
	// 	height: 2.24,
	// 	width:  2.24,
	// }

	// printArea(areaRectangle)
	sq := square{2}

	printArea(sq)
}
