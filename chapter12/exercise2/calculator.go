// Calculator is a simple RPN-style Calculator program that receives integers
// up to 999999, and performs operations (+, -, *, /)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./stack"
)

func main() {
	s := stack.New()
	rd := bufio.NewReader(os.Stdin)
	fmt.Println("Enter an integer up to 999999, or and operator (+, -, *, /), then enter.")

	for {
		fmt.Println(s)
		// if s.Index == 0 {
		// 	fmt.Printf("\t y: 0\n")
		// 	fmt.Printf("x: 0\n")
		// 	fmt.Println()
		// } else if s.Index == 1 {
		// 	fmt.Printf("\t y: 0\n")
		// 	fmt.Printf("x: %v\n", s.Value[0])
		// 	fmt.Println()
		// } else {
		// 	fmt.Printf("\t y: %v\n", s.Value[1])
		// 	fmt.Printf("x: %v\n", s.Value[0])
		// 	fmt.Println()
		// }
		fmt.Printf("\t y: %v\n", s.Value[1])
		fmt.Printf("x: %v\n", s.Value[0])
		fmt.Println()

		input, err := rd.ReadString('\n')
		input = strings.Trim(input, "\n")
		if err != nil {
			fmt.Println("SOMETHING BAD HAPPENED!!")
			break
		}

		switch {
		case input == "q":
			return
		case input == "+":
			nm := (s.Value[1] + s.Value[0])
			s.Pull()
			s.Value[0] = nm
			s.Delete()
		case input == "-":
			nm := (s.Value[1] - s.Value[0])
			s.Pull()
			s.Value[0] = nm
			s.Delete()
		case input == "*":
			nm := (s.Value[1] * s.Value[0])
			s.Pull()
			s.Value[0] = nm
			s.Delete()
		case input == "/":
			nm := (s.Value[1] / s.Value[0])
			s.Pull()
			s.Value[0] = nm
			s.Delete()
		case input < "999999":
			num, _ := strconv.Atoi(input)
			s.Push(num)
		default:
			fmt.Println("I said enter a number or operator!!")
		}
	}
}
