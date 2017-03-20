// Function fibonacci
package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		num, step := fibonacci(i)
		fmt.Printf("Fibonacci num %d is: %d", step, num)
	}
}

func fibonacci(n int) (res, step int) {
	step = n
	if n <= 0 {
		res = 1
	} else {
		fib1, _ := fibonacci(n - 1)
		fib2, _ := fibonacci(n - 2)
		res = fib1 + fib2
	}
	return
}
