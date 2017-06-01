package main

import (
	"fmt"

	"./min_interface"
)

func main() {
	// intarray := min.IntArray{45, 85, 65, 35, 25, 15, 45, 75, 96}
	stringarray := min.StringArray{"aa", "savbksb", "sbsbisi", "zgheoenh", "ghr0qppwidn"}

	// fmt.Println(min.Min(intarray))
	fmt.Println(min.Min(stringarray))
}
