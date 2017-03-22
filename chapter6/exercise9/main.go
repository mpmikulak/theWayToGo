// A non-recursive fibonacci calculator
package main

import (
	"fmt"
	"time"
)

func main() {
	var result int

	fib := func(v int) {
		for i := 0; i <= v; i++ {
			result += i
		}
	}
	timeBefore := time.Now()
	fib(1e8)
	fmt.Println(result)
	timeAfter := time.Now()
	delta := timeAfter.Sub(timeBefore)
	fmt.Println(delta)
}
