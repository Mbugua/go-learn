package main

import "fmt"

// common interface that defines getArea of a shape
type shape interface {
	getArea() float64
}

// custom struct for triangle dimension with
// height of type float64
// and base of type float 64
type triangle struct {
	height float64
	base   float64
}

// custom struct for square dimension with
// sideLength of type float64
type square struct {
	sideLength float64
}

// function to get area of a triangular shape.
// Area of a triangle = 0.5 * base * height
func (t triangle) getArea() float64 {
	tArea := 0.5 * t.base * t.height

	return tArea
}

// function to get area of a square shape.
// Area of square = sideLength * sideLength
func (sq square) getArea() float64 {
	sqArea := sq.sideLength * sq.sideLength

	return sqArea
}

// function to calculate and print area of shape
func printArea(s shape) {
	fmt.Println(s.getArea())
}

// main function to run program
func main() {
	t := triangle{
		height: 2.4,
		base:   4.2,
	}

	// calcuate area of triangular shape
	fmt.Println("Area of the triangle is:")
	printArea(t)

	sq := square{
		sideLength: 2.5,
	}
	// calculate area of sqaure
	fmt.Println("Area of the square is:")
	printArea(sq)
}
