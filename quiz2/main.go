package main

import (
	"fmt"
)

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	tri := triangle{
		base:   2,
		height: 4,
	}
	sqr := square{
		sideLength: 3,
	}

	printArea(tri)
	printArea(sqr)
}

func printArea(s shape) {
	fmt.Println("Area is :", s.getArea())
}

func (t triangle) getArea() float64 {
	area := t.base * t.height * 0.5
	return area
}

func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}
