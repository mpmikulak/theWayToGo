// Defines function Filter that accepts []int and another function, and
// returns a slice of only the even numbers
package main

import "fmt"

func main() {
	slice := make([]int, 10)
	for i := 1; i <= len(slice); i++ {
		slice[i-1] = i
	}
	fmt.Println(slice)
	slice = Filter(slice, isEven)
	fmt.Println(slice)
}

func Filter(s []int, fn func(int) bool) []int {
	ret := make([]int, 0)
	for _, n := range s {
		if fn(n) {
			ret = append(ret, n)
		}
	}
	return ret
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
