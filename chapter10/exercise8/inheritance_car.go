// Make a working example for the Car and Engine code above; also give the Car
// type a field wheelCount and a method numberOfWheels() to retrieve this. Make
// a type Mercedes which embeds Car, an object of type Mercedes and use the
// methods. Then construct a method sayHiToMerkel() only on type Mercedes and
// invoke it.
package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
	Engine
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi, Merkel.")
}

func (c *Car) numberOfWheels() int {
	return c.wheelCount
}

func (c *Car) GoToWorkIn() {
	c.Start()
	c.Stop()
}

func main() {
	myNewCar := new(Mercedes)
	myNewCar.wheelCount = 4
	fmt.Println(myNewCar.numberOfWheels())
	myNewCar.sayHiToMerkel()
}
