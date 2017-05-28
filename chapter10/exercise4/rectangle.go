// Define a struct Rectangle with int properties length and width. Give this
// type methods Area() and Perimeter() and test it out.
package main

import "fmt"

type Rectangle struct {
	length    float64
	width     float64
	area      float64
	perimeter float64
}

func (r *Rectangle) Area() {
	r.area = r.width * r.length
}

func (r *Rectangle) Perimeter() {
	r.perimeter = r.length*2 + r.width*2
}

func main() {
	rect := new(Rectangle)
	rect.length = 8
	rect.width = 9
	fmt.Println(rect)

	rect.Area()
	fmt.Println(rect)

	rect.Perimeter()
	fmt.Println(rect)
}
