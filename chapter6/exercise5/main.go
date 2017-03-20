// Prints the numbers from 10 to 1 in that order using recursive functions
package main

import "fmt"

func main() {
	printrec(10)
}

func printrec(i int) {
	if i <= 0 {
		return
	}
	fmt.Println(i)
	printrec(i - 1)
}
