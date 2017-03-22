// Proves that when using new, the return value is a pointer to the
// array and it must  be dereferenced
package main

import "fmt"

func main() {
	var arr1 = new([10]int)
	arr2 := *arr1 // Dereferencing the array passes a copy of the values
	arr2[0] = 1

	fmt.Println(arr1)
	fmt.Println(arr1)

}
