// Define a struct employee with a field salary, and make a method giveRaise
// for this type to increase the salary with a certain percentage.
package main

import "fmt"

type employee struct {
	salary float32
}

func (e *employee) giveRaise(p float32) {
	if p < 1 {
		e.salary += e.salary * p
	} else {
		e.salary += e.salary * p / 100.0
	}
}

func main() {
	mike := employee{100}
	fmt.Println(mike.salary)

	mike.giveRaise(2.5) // Entered as whole percent
	fmt.Println(mike.salary)

	mike2 := employee{100}
	mike2.giveRaise(.025) // Entered as fractional
	fmt.Println(mike2.salary)
}
