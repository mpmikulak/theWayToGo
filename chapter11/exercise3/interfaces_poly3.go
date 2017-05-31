package main

import (
	"fmt"
	"math"
)

// Contract interface
type Shaper interface {
	Area() float32
}

// Empty shape type and area method
type Shape struct{}

func (sh Shape) Area() float32 {
	return 0
}

// Square type and area method
type Square struct {
	side float32
}

func (sq *Square) Area() float32 {

	return sq.side * sq.side
}

// Rectangle type and method
type Rectangle struct {
	length, width float32
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

// Circle type and method
type Circle struct {
	radius float32
}

func (c *Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r := &Rectangle{8, 7}
	s := &Square{5}
	c := &Circle{6}
	sh := Shape{}
	shapes := []Shaper{r, s, c, sh}
	for i := range shapes {
		fmt.Println(shapes[i].Area())
	}
}
