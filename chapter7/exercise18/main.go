// Defines a map function that takes a lambda function and an array of ints
// and returns the items in the array multiplied by 10
package main

import "fmt"

func main() {
	fun := func(n int) int { return n * 10 }
	list := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list = mapFunc(fun, list)
	fmt.Println(list)
}

func mapFunc(f func(int) int, l [10]int) (r [10]int) {
	for i := 0; i < len(l); i++ {
		r[i] = f(l[i])
	}
	return
}
