package main

import "fmt"

func main() {
	arr := new([16]int)
	// fmt.Println(*arr)
	fmt.Printf("%v\t%T\n", *arr, *arr)
	fmt.Printf("%v\t%T\n", arr, arr)
	for i := range arr {
		arr[i] = i
	}
	fmt.Println(*arr)
}
