package main

import (
	"fmt"
	"math"
)

// INTERFACES
type Magnitude interface {
	Abs() float64
}

// 2 Dimensional point type and method
type Point struct {
	x int
	y int
}

func (p *Point) Abs() float64 {
	return math.Sqrt(float64(p.x*p.x) + float64(p.y*p.y))
}

func (p *Point) Scale(s int) {
	p.x = p.x * s
	p.y = p.y * s
}

// Three dimensional type and method
type Point3 struct {
	x int
	y int
	z int
}

func (p *Point3) Abs() float64 {
	twoDimension := math.Sqrt(float64(p.x*p.x) + float64(p.y*p.y))
	return math.Sqrt(float64(p.z*p.z) + float64(twoDimension*twoDimension))
}

func (p *Point3) Scale(s int) {
	p.x = p.x * s
	p.y = p.y * s
	p.z = p.z * s
}

// Polar type and method
type Polar struct {
	Distance float64
	Angle    float64
}

func (p *Polar) Abs() float64 {
	return p.Distance
}

func main() {
	// Initialize the point types
	pt := &Point{8, 9}
	pt3 := &Point3{1, 2, 3}
	pol := &Polar{7.8, 45}

	var wreckingBall Magnitude

	wreckingBall = pt
	fmt.Println("The magnitude of the point is: ", wreckingBall.Abs())

	wreckingBall = pt3
	fmt.Println("The magnitude of the point is: ", wreckingBall.Abs())

	wreckingBall = pol
	fmt.Println("The magnitude of the point is: ", wreckingBall.Abs())
}
