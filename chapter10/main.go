package main

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type myStruct struct{ i int }

type Interval struct {
	start int
	end   int
}

func main() {

	// Allocating memory for a struct
	ms := new(struct1) // Create a new variable with type struct1
	ms.i1 = 10         // Assign values to the diferent fields of the variable
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Println(ms)

	// Using a struct literal
	mq := &struct1{10, 15.5, "Chris"} // Idiomatic way of initializing the above vaiable
	fmt.Println(mq)
	var mt struct1

	mt = struct1{10, 15.5, "Chris"} // Another way of doing it idiomatically
	fmt.Println(mt)

	// Some more initializations
	inter := Interval{0, 3}              // Fields must be in order
	inter2 := Interval{end: 5, start: 1} // Field names precede the value
	inter3 := Interval{end: 5}           // A field is ommited
}
