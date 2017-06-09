package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Enter a name!!")
		return
	}

	for i := 1; i < len(os.Args); i++ {
		r := os.Args[i]
		fmt.Printf("Hello, %v!\n", r)
	}
}
