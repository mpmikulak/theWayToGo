package main

import (
	"fmt"

	"./fibo"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci number %v is %v\n", i, fibo.Fibonacci(i))
		fmt.Println("The last input number is: ", fibo.Result)
	}
}
