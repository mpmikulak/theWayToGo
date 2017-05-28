// Define a struct type Base which contains an id-field, and methods Id() to
// return the id and SetId() to change the id. The struct type Person contains
// Base, and fields FirstName and LastName. The struct type Employee contains a
// Person and a salary. Make an employee instance and show its id.
package main

import "fmt"

type Base struct {
	id int
}

type Person struct {
	Base
	FirstName string
	LastName  string
}

type Employee struct {
	Person
	salary float64
}

func (b *Base) Id() int {
	return b.id
}

func (b *Base) SetId(i int) {
	b.id = i
}

func main() {
	worker := &Employee{Person{Base{123456}, "Billy", "Idol"}, 1500.90}
	fmt.Println(worker.Id())

}
