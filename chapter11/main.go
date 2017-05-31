package main

import (
	"fmt"
	"math"
)

// Shaper interface allows both squares and Rectangles to be grouped together
type Shaper interface {
	Area() float32
}

type Circle struct {
	radius float32
}

func (c *Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

type Square struct {
	side float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(v valuable) {
	fmt.Printf("The value of the asset is $%.2f\n", v.getValue())
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("param #%d is an int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d’s type is unknown\n", i)
		}
	}
}

func main() {
	r := Rectangle{5, 3}
	q := &Square{5}

	shapes := []Shaper{r, q}
	fmt.Println("Loopint through shapes for area...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)

	// Type assertions
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1

	if t, ok := areaIntf.(*Square); ok { // Type assertions return the type and whether it matches the asserted type
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if t, ok := areaIntf.(*Circle); ok { // Type assertions return the type and whether it matches the asserted type
		fmt.Printf("The type of areaIntf is: %T\n", t)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

	// THE TYPE SWITCH!!!!!
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type circle %T with value %v\n", t, t)
	// case float32:
	// 	fmt.Printf("Type float32 %T with value %v\n", t, t)
	case nil:
		fmt.Println("Nil value")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}
