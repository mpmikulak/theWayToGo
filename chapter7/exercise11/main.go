// Defines a function that inserts a string slice into another string
// slice at a given location
package main

import "fmt"

func main() {
	slice := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	insertSlice := []string{"h", "e"}
	limit := len(slice)
	for i := 0; i <= limit; i++ {
		res := InsertStringSlice(slice, insertSlice, i)
		fmt.Println(res)
	}
}

func InsertStringSlice(t []string, i []string, l int) []string {
	newSlice := make([]string, len(t)+len(i))
	copy(newSlice, t[:l])
	copy(newSlice[l:], i)
	totalIndex := l + len(i)
	fmt.Println(totalIndex)
	copy(newSlice[totalIndex:], t[l:])
	return newSlice
}
