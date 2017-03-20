package main

import (
	"fmt"
	"time"
)

func main() {
	timeBefore := time.Now()
	func() {
		sum := 0
		for i := 1; i <= 1e8; i++ {
			sum += i
		}
		fmt.Println(sum)
	}()
	timeAfter := time.Now()

	seconds := timeAfter.Second() - timeBefore.Second()
	nanoseconds := timeAfter.Nanosecond() - timeBefore.Nanosecond()

	fmt.Printf("%v : %v\n", seconds, nanoseconds)
}
