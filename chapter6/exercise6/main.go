// Calculates the factorial of a given number
package main

import (
	"fmt"
	"math/big"
)

func main() {
	for i := 1; i <= 10000; i++ {
		fmt.Printf("Factorial %d is: %v\n", i, factorial(i))
	}
}

func factorial(n int) (x *big.Int) {
	one := big.NewInt(1)
	add := big.NewInt(0)
	result := big.NewInt(1)

	for i := n; i > 0; i-- {
		result.Mul(result, add.Add(add, one))
	}
	return result
}

// func factorial_2(n int) (f int){
//
// }
