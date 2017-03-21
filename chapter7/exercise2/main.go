package main

import "fmt"

func main() {
	arr := new([16]int)
	fmt.Println(*arr)
	for i := range arr {
		arr[i] = i
	}
	fmt.Println(*arr)
}
