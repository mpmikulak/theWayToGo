// Summary of chapter 8 - Maps
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Maps are reference types with key value pairs
	// Any value that can be checked by operations == and != can
	// be used as the key, including pointers and interface types
	// Maps can be created like any other literal
	mapLit := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Printf("Map literal: %v , type: %T\n", mapLit, mapLit)

	// Also using the make keyword. Do not use the new keyword
	mapCreated := make(map[string]float32)
	mapCreated["half"] = 0.5
	fmt.Printf("Map created: %v , type: %T\n", mapCreated, mapCreated)

	// Maps can be assigned as well
	mapAssigned := mapCreated
	fmt.Printf("Map assigned: %v , type: %T\n", mapAssigned, mapAssigned)

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
	items := make([]map[int]int, 5) // Specify the slice to be 5 long
	for i := range items {
		items[i] = make(map[int]int) // The index of the map must be used
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	items2 := make([]map[int]int, 5)
	for _, item := range items2 { // "item" is a copy of the value
		item = make(map[int]int, 1)
		item[1] = 2 // This doesn't change the value
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)

	// Maps aren't sorted, even the keys
	// To sort a map, copy to a slice and sort using the sort package
	var (
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16,
			"indio": 87, "juliet": 65, "kilo": 43, "lima": 98}
	)

	fmt.Println("Unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	fmt.Println()

	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("Sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
	fmt.Println()
}
