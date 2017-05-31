// Define an interface Simpler with methods Get() which returns an integer,
// and Set() which has an integer as parameter. Make a struct type Simple
// which implements this interface. Then define a function which takes a
// parameter of the type Simpler and calls both methods upon it. Call this
// function from main to see if it all works correctly.
package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	number int
}

func (s *Simple) Get() int {
	return s.number
}

func (s *Simple) Set(i int) {
	s.number = i
}

func call(n int, s Simpler) int {
	s.Set(n)
	return s.Get()
}

func main() {
	sim := &Simple{5}
	call(7, sim)
	fmt.Println(sim)
}
