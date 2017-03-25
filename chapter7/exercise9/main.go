// Defines a function that enlarges a given []int by a factor
package main

import "fmt"

func main() {
	slice := []int{5, 6, 7, 8, 2, 4, 5, 6}
	fmt.Printf("%v\t%v\n", len(slice), cap(slice))
	slice = magnifySlice(slice, 2)
	fmt.Printf("%v\t%v\n", len(slice), cap(slice))
}

func magnifySlice(s []int, f int) []int {
	length := len(s)
	newSlice := make([]int, length*f)
	copy(newSlice, s)
	s = newSlice
	return s
}
