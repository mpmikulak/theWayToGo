// Proves that when assigning an array to another, a distinct copy
// in memory is made
package main

import "fmt"

func main() {
	array1 := new([5]int)
	array2 := array1
	array2[0] = 1

	fmt.Println(&array1)
	fmt.Println(&array2)

}
