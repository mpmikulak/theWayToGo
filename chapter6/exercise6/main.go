// Calculates the factorial of a given number using arbitrary precision
package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i <= 1000; i++ {
		fmt.Printf("Factorial %d is: %v\n", i, factorial(i))
	}
	finish := time.Now()
	delta := finish.Sub(start)
	fmt.Printf("Took %v\n", delta)

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
