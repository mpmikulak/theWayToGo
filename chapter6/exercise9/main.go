// A non-recursive fibonacci calculator
package main

import "fmt"

func main() {
	var result int

	fib := func(v int) {
		for i := 0; i <= v; i++ {
			result += i
		}
	}
	fib(1500)
	fmt.Println(result)
}
