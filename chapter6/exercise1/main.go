package main

import "fmt"

var a, b int = 4, 2

func main() {

	c, d, e := returnVal(a, b)
	fmt.Printf("Sum: %v, Product: %v, Difference: %v.\n", c, d, e)

	c, d, e = returnVal_2(a, b)
	fmt.Printf("Sum: %v, Product: %v, Difference: %v.\n", c, d, e)

} /////////////////////////////////////

// Function with named return values
func returnVal(x, y int) (sum, prod, diff int) {
	sum = x + y
	prod = x * y
	diff = x - y
	return
}

// Function with unnamed return values
func returnVal_2(x, y int) (int, int, int) {
	return x + y, x * y, x - y
}
