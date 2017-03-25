package main

import (
	"fmt"
	"time"
)

const limit = 51

func main() {

	arr := new([limit]int)
	timeBefore := time.Now()
	arr[0] = 0
	arr[1] = 1

	for i := 2; i < limit; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	timeAfter := time.Now()

	delta := timeAfter.Sub(timeBefore)
	fmt.Printf("Fibonacci number %v: %v. Calculated in %v.\n", len(arr), arr[len(arr)-1], delta)
}
