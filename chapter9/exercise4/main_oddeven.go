package main

import (
	"fmt"

	"./even"
)

func main() {
	for i := 0; i < 100; i++ {
		if even.IsEven(i) {
			fmt.Println(i)
		}
	}
}
