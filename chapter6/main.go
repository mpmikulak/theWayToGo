package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {

	// A simple example of a function calling another function
	println("In main before calling a greeting")
	greeting()
	println("In main after calling a greeting")

	// A function can take another function as an argument, provided
	// it returns the apropriate values
	fmt.Println(takeInts(returnInts()))

	// A function can also be given as a type. They are first class values
	// type binOp func(int, int) int
	// var add binOp
	// println(add(1, 2))

	// Named return values vs unnamed
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()

	// A function with a variadic parameter can give this parameter
	// to different functions
	str1 := "This is a string"
	str2 := "This is a second string"
	F1(str1, str2)

	// Multiple defer statements execute in reverse order
	F4()

	// New vs make:
	// New allocates zeroed storage for a new item of type T and
	// returns its address (a pointer). Used for value types
	v := new(int) // Has type *int
	fmt.Println(v)
	fmt.Println(*v)

	// Make returns an initialized variable of type T, so it does more
	// work than new.  Used for built-in reference types
	t := new(map[string]int)
	fmt.Println(t)

	// Recursive functions
	result := 0
	step := 0
	for i := 0; i <= 10; i++ {
		result, step = fibonacci(i)
		fmt.Printf("Fibonaci %d is: %d\n", step, result)
	}

	callback(1, Add)

	// Lambda functions
	func() {
		sum := 0
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
	}()

	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 3:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))

} /////// END of MAIN

func greeting() {
	println("I'm a greeting: Hi!!")
}

func returnInts() (int, int, int) {
	return 1, 2, 3
}

func takeInts(a, b, c int) int {
	return a + b + c
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

// Unnamed return: less good
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// Named return: more good
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}

func F1(s ...string) {
	F2(s...)
	F3(s)
}

func F2(s ...string) {
	for _, i := range s {
		fmt.Println(i)
	}
}

func F3(s []string) {
	for _, i := range s {
		fmt.Println(i)
	}
}

// A demonstartion of defer's LIFO
func F4() {
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
}

func fibonacci(n int) (res, step int) {
	step = n
	if n <= 0 {
		res = 1
	} else {
		fib1, _ := fibonacci(n - 1)
		fib2, _ := fibonacci(n - 2)
		res = fib1 + fib2
	}
	return
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}
func callback(y int, f func(int, int)) {
	f(y, 2)
	// this becomes Add(1, 2)
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
