package main

import (
	"fmt"

	"./sorting"
)

// Person definition
type Person struct {
	firstName string
	lastName  string
}

func (p Person) String() string {
	return p.firstName + " " + p.lastName
}

// Persons definition and methods
type Persons []Person

func (p Persons) Len() int      { return len(p) }
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Persons) Less(i, j int) bool {
	f := p[i].firstName + " " + p[i].lastName
	s := p[j].firstName + " " + p[j].lastName
	return f < s
}

func main() {
	p1 := Person{"James", "Dean"}
	p2 := Person{"Michael", "Jordan"}
	p3 := Person{"Randy", "Johnson"}

	company := Persons{p1, p2, p3}
	fmt.Println(company)
	sorting.Sort(company)
	fmt.Println(company)
}
