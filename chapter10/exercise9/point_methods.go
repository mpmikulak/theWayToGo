// Define a 2 dimensional Point with coordinates X and Y as a struct. Do the
// same for a 3 dimensional point, and a Polar point defined with its polar
// coordinates. Implement a function Abs() that calculate the length of the
// vector represented by a Point, and a function Scale that multiplies the
// coordinates of a point with a scale factor(hint: use function Sqrt from
// package math)
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x      int
	y      int
	vector float64
}

type ThreePoint struct {
	x      int
	y      int
	z      int
	vector float64
}

type Polar struct {
	Distance float64
	Angle    float64
}

func (p *ThreePoint) ThreePointAbs() {
	twoDimension := math.Sqrt(float64(p.x*p.x) + float64(p.y*p.y))
	p.vector = math.Sqrt(float64(p.z*p.z) + float64(twoDimension*twoDimension))
}

func (p *Point) Abs() {
	p.vector = math.Sqrt(float64(p.x*p.x) + float64(p.y*p.y))
}

func (p *Point) Scale(s int) {
	p.x = p.x * s
	p.y = p.y * s
}

func main() {
	// Two point calculations
	coord := new(Point)
	coord.x = 5
	coord.y = 7

	coord.Abs()
	fmt.Printf("%v %T\n", coord, coord)

	coord.Scale(7)
	coord.Abs()
	fmt.Printf("%v %T\n", coord, coord)

	// Three point calculations
	threeD := new(ThreePoint)
	threeD.x = 7
	threeD.z = 12
	threeD.y = 6
	fmt.Println(threeD)

	threeD.ThreePointAbs()
	fmt.Println(threeD)
}
