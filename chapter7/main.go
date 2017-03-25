// Summary of chapter 7
package main

import "fmt"

func main() {
	// Slices:

	// Slices are delcared by var, the name of the slice, then the type
	// with no size
	// var sliceName []int

	// Format for the initialization of a slice:
	// var slice1 []type = arr1[start:end] <- this is a slice-expression
	// slice1 []type = arr1[:] == slice1 = &arr1

	// A slice literal
	s := []int{1, 2, 3}
	fmt.Println(cap(s))
	// A second slice but refering to the same array
	// s2 := s[1:2]
	// A slice can be expanded to its maximum size with:
	s = s[:cap(s)]
	fmt.Println(cap(s))
	// fmt.Println(cap(s2))

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Printf("%c\n", b[1:4])
	fmt.Printf("%c\n", b[:2])
	fmt.Printf("%c\n", b[2:])
	fmt.Printf("%c\n", b[:])

	// Creating a slice with make
	// s2 = make([]int, 10)

	// A slice of bytes can be made out of a string immediately like this
	str := "golang string"
	c := []byte(str)
	fmt.Printf("%c\n", c)

	r := []rune(str)
	fmt.Printf("%c\n", r)
}
