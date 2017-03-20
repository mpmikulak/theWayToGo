// Main contains a Lambda that prints "Hello, World" and then assigns it
// to a variable, then checks the type of the variable
package main

import "fmt"

func main() {
	world := func() { fmt.Println("Hello, World!!") }
	world()
	fmt.Printf("%T\n", world)
}
