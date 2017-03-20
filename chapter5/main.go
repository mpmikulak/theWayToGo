package main

import "fmt"

func main() {

	// Typical if statement
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%5 == 0 {
			fmt.Println("Fizz")
		} else if i%3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	for i := 0; i <= 100; i++ {
		modulus := i % 3

		// Demonstration of a value switch
		switch modulus { // Variable is evaluated once
		case 1:
			fmt.Println("ONE")
		case 2:
			fmt.Println("TWO")
		case 3:
			fmt.Println("THREE")
		case 4:
			fmt.Println("FOUR")
		case 5:
			fmt.Println("FIVE")
		case 6:
			fmt.Println("SIX")
		case 7:
			fmt.Println("SEVEN")
		case 8:
			fmt.Println("EIGHT")
		case 9:
			fmt.Println("NINE")
		default:
			fmt.Println("ZERO!")
		}
	}
	fmt.Println()

	// A switch true statement
	for i := 0; i <= 100; i++ {
		switch { // Each case is evaluated separately
		case i%15 == 0:
			fmt.Printf("%v: Fizz Buzz\n", i)
		case i%5 == 0:
			fmt.Printf("%v: Fizz\n", i)
		case i%3 == 0:
			fmt.Printf("%v: Buzz\n", i)
		default:
			fmt.Println(i)
		}
	}
	fmt.Println()

	// Demonstration of fallthrough
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

	// Demonstration of using a label and continue
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1 // Takes the control flow to LABEL1
			}
			fmt.Printf("I is : %d, and j is: %d\n", i, j)
		}
	}

	// USing a continue statement to jump back to the beginning of the loop
	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}

	// Using a goto statement
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE

}
