package main

import (
	"fmt"
)

// INTERFACES
type AreaInterface interface {
	Area() float32
}

type PeriInterface interface {
	Perimeter() float32
}

// Square type and methods
type Square struct {
	side float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func (s *Square) Perimeter() float32 {
	return s.side * 4
}

// Triangle type and methods
type Triangle struct {
	height float32
	width  float32
}

func (t *Triangle) Area() float32 {
	return t.height * t.width / 2
}

func main() {
	var areaIntf AreaInterface
	var periIntf PeriInterface

	sq := &Square{7.5}
	tr := &Triangle{9.2, 5.7}

	areaIntf = sq
	fmt.Println("The area of the square is: ", areaIntf.Area())

	areaIntf = tr
	fmt.Println("The area of the triangle is: ", areaIntf.Area())

	periIntf = sq
	fmt.Println("The perimeter of the square is: ", periIntf.Perimeter())
}
