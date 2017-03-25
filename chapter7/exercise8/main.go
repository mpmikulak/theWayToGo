package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Generate a slice of random numbers
	nums := make([]int, 30)
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(100)
	}
	// Print the numbers before min/max operations
	fmt.Println(nums)
	// Pass the slice to the functions
	max := maxSlice(nums)
	min := minSlice(nums)
	fmt.Println(max)
	fmt.Println(min)
}

func minSlice(slice []int) int {
	min := slice[0]
	for _, i := range slice {
		if i < min {
			min = i
		}
	}
	return min
}

func maxSlice(slice []int) int {
	max := slice[0]
	for _, i := range slice {
		if i > max {
			max = i
		}
	}
	return max
}
