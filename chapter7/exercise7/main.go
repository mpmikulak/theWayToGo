package main

import "fmt"

func main() {
	array := [4]float64{5.5, 6.6, 7.7, 8.8}

	sum := Sum(array)
	fmt.Println(sum)

	slice := array[:]
	sum = Sum2(slice)
	fmt.Println(sum)

	ints := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(SumAndAverage(ints))
}

// Sum function taking an array and summing the contents
func Sum(arr [4]float64) (total float64) {
	for _, num := range arr {
		total += num
	}
	return
}

// Sum function taking a slice and summing the contents
func Sum2(slice []float64) (total float64) {
	for _, num := range slice {
		total += num
	}
	return
}

// Function that takes a slice of ints and returns the sum and the average
func SumAndAverage(slice []int) (int, float32) {
	sum := 0
	for _, i := range slice {
		sum += i
	}
	average := (float32(sum) / float32(len(slice)))
	return sum, average
}
