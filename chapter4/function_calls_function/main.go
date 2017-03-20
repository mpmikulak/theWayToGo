package main

import "fmt"

var a string

func main() {
	a = "G"
	print(a)
	f1()
	fmt.Println()

	five := 5
	// six := 6
	num1 := 5
	// num2 := 6

	bool1 := (five == num1)
	// bool2 := (six ==)
	fmt.Println(bool1)
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}
