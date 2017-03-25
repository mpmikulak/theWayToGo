// Summary of chapter 8 - Maps
package main

import "fmt"

func main() {
	// Maps are reference types with key value pairs
	// Maps can be created like any other literal
	mapLit := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(mapLit)

	// Also using the make keyword. Do not use the new keyword
	mapCreated := make(map[string]float32)
	mapCreated["half"] = 0.5
	fmt.Println(mapCreated)

	// Maps can be assigned as well
	mapAssigned := mapCreated
	fmt.Println(mapAssigned)

	// Maps can have slices as values
	sliceMap := make(map[int][]int)
	sliceMap[1] = []int{1, 2, 3, 4, 5}
	fmt.Println(sliceMap)
	fmt.Println(sliceMap[1])

	// If a key is used against a map but the value doesn't exist, it is
	// initialized to it's zero value. Use the comma ok syntax to check.
	// The map will return true or false depending on whether the value
	// exists or not
	_, ok := mapCreated["five"]
	fmt.Println(ok)

	// The for range construct works on maps also
	// for key, value := range map1
	// For both the key and value:
	for key, value := range mapLit {
		fmt.Printf("%v\t%v\n", key, value)
	}

	// For just the value:
	for _, value := range mapLit {
		fmt.Println(value)
	}

	// For just the key:
	for key := range mapLit {
		fmt.Println(key)
	}

	// Use make to create a slice of maps
	mapSlice := make([]map[int]int, 5) // Specify the slice to be 5 long
	for i := range mapSlice {
		mapSlice[i] = make(map[int]int)
		mapSlice[i][1] = 2
	}
	fmt.Println(a)
}
