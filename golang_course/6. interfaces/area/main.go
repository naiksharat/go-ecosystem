package main

import "fmt"

type triangle struct {
	height float64
	width  float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.width * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {

	t := triangle{1, 2}
	s := square{1}

	printArea(t)
	printArea(s)
}
