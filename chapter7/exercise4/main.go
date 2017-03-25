// Fibonacci function that returns a slice
package main

import "fmt"

func main() {
	slice := fibonacci(10)
	fmt.Println(slice)
}

func fibonacci(n int) (s []int) {
	s = make([]int, n+1)
	s[0] = 0
	s[1] = 1
	for i := 2; i < len(s); i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return
}
