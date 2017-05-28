package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Integer int

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    // Anonymous field
	innerS // Imbedded struct
}

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

type Person struct {
	firstName string "This is a tag" // Can only be accessed by the reflect package
	lastName  string "More tags"
}

func (p Integer) get() int { return int(p) }

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

type File struct { // Used for the factory below
	fd   int    // File descriptor number
	name string // File name
}

func NewFile(fd int, name string) *File { // Returns a pointer to a new struct
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func refTag(tt Person, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

var myInteger Integer

func main() {

	// Allocating memory for a struct
	ms := new(struct1) // Create a new variable with type struct1
	ms.i1 = 10         // Assign values to the diferent fields of the variable
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Println(ms)

	// Using a struct literal
	pointerStruct := &struct1{10, 15.5, "Chris"} // Idiomatic way of initializing the above vaiable
	normalStruct := struct1{20, 155.5, "James"}
	fmt.Println(pointerStruct)
	fmt.Println(normalStruct)
	var mt struct1

	mt = struct1{10, 15.5, "Chris"} // Another way of doing it idiomatically
	fmt.Println(mt)

	// Some more initializations
	inter := Interval{0, 3}              // Fields must be in order
	inter2 := Interval{end: 5, start: 1} // Field names precede the value
	inter3 := Interval{end: 5}           // A field is ommited
	fmt.Printf("%v %v %v \n", inter, inter2, inter3)

	// Three ways to call a method with a user defined struct
	// #1 - struct as a value type:
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %v %v\n", pers1.firstName, pers1.lastName)

	// #2 - struct as a pointer:
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	upPerson(pers2)
	fmt.Printf("The name of the person is %v %v\n", pers2.firstName, pers2.lastName)

	// #3 - struct as a literal
	pers3 := &Person{"Chris", "Woodward"} // Returns a pointer
	upPerson(pers3)
	fmt.Printf("%v\n", pers3)
	fmt.Printf("The name of the person is %v %v\n", pers3.firstName, pers3.lastName)

	// No constuctors in Go but constuctor-like factory functions. Convention
	// is that the function name starts with new. Factory functions
	// "instantiate" an object of a defined type.
	f := NewFile(10, "./test.txt")
	fmt.Println(f)

	// Index through the tags of a struct
	for i := 0; i < 2; i++ {
		refTag(pers1, i)
	}

	// Using anonymous fields in a struct
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// With a struct literal
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)

	myInteger = 7
	fmt.Printf("%v %T\n", myInteger, myInteger)

	myInt := myInteger.get()
	fmt.Printf("%v %T\n", myInt, myInt)

	// To display the current memory status
	memory := new(runtime.MemStats)
	runtime.ReadMemStats(memory)
	fmt.Printf("%v\n", memory.Alloc/1024)
}
