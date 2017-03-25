// Function fibonacci
package main

import (
	"fmt"
	"time"
)

func main() {

	timeBefore := time.Now()
	num, step := fibonacci(50)
	timeAfter := time.Now()
	delta := timeAfter.Sub(timeBefore)
	fmt.Printf("Fibonacci num %d is: %d. Calculated in %v\n", step, num, delta)

	// for i := 0; i <= 10; i++ {
	// }
}

func fibonacci(n int) (res, step int) {
	step = n
	if n == 0 {
		res = 0
		return
	} else if n == 1 {
		res = 1
	} else if n == 2 {
		res = 1
	} else {
		fib1, _ := fibonacci(n - 1)
		fib2, _ := fibonacci(n - 2)
		res = fib1 + fib2
	}
	return
}
