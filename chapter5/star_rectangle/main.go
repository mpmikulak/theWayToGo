// Prints a stupid box made out of stars
package main

import "fmt"

func main() {

	for height := 10; height != 0; height-- {

		for width := 20; width != 0; width-- {
			print("*")
		}

		println()
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d", v)
		v = 5
	}

	for i := 0; ; i++ {
		fmt.Println("The value of i is now:", i)
	}
}
