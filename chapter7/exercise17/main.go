// Sort a slice of ints through a function using the bubble-sort method
package main

import "fmt"

func main() {
	ints := []int{5, 6, 8, 7, 2, 3, 1, 1, 5, 8, 9, 2, 1, 87, 36}
	ints = bubbleSort(ints)
	fmt.Println(ints)

}
func bubbleSort(s []int) []int {
	for i := len(s); i > 0; i-- {
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
	}
	return s
}
