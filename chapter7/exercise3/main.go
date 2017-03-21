package main

import (
	"fmt"
	"time"
)

const limit = 50

func main() {
	timeBefore := time.Now()
	result := fibonacci(limit)
	timeAfter := time.Now()
	delta := timeAfter.Sub(timeBefore)
	fmt.Printf("Fibonacci number %v calculated in %v.\n", limit, delta)
	fmt.Println(result)
}

func fibonacci(x int) (fib int) {
	fibArr := new([limit]int)
	fibArr[0] = 1
	fibArr[1] = 2
	for i := 2; i < limit; i++ {
		fibArr[i] = fibArr[i-1] + fibArr[i-2]
	}
	fib = fibArr[limit-1]
	return
}
