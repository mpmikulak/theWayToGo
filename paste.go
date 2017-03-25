package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10)
	fmt.Println(len(s))
	s = s[1 : 1+1]
	fmt.Println(len(s))
}
