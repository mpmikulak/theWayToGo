package main

import "fmt"

func main() {
	loopThrough(1, 2, 3, 4, 5, 6, 7, 7)
}

func loopThrough(a ...int) {
	for _, i := range a {
		fmt.Println(i)
	}
}
