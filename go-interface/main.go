package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

// define a interface
type shape interface {
	getArea() float64
}

func main() {
	myTriangle := triangle{
		height: 2,
		base:   3,
	}

	mySquare := square{
		sideLength: 3,
	}

	printArea(myTriangle)
	printArea(mySquare)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
