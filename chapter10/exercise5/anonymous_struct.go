// Make a struct with 1 named float field, and 2 anonymous fields of type int and
// string. Create an instance of the struct with a literal expression and print
// the content.
package main

import "fmt"

type anon struct {
	floating float32
	int
	string
}

func main() {
	anonymous := anon{3.14, 1222220254, "Stringing"}
	fmt.Println(anonymous)
}
