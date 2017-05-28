package main

import (
	"fmt"
	"strconv"
)

const LIMIT = 4

type Stack [LIMIT]int

// Push finds the highest cell with value 0 and places the argument there
func (s *Stack) Push(n int) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 0 {
			s[i] = n
			return
		}
	}
	fmt.Println("OH NOES, AWW FUL!?")
	return
}

// Pop finds the lowest cell thats not empty and returns the value after clearing that cell
func (s *Stack) Pop() int {
	for ix, v := range s {
		if v != 0 {
			s[ix] = 0
			return v
		}
	}
	return 0
}

// String returns the formatted string with the current contents of the stack
func (s *Stack) String() string {
	str := ""
	for ix, v := range s {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(v) + "]"
	}
	return str
}

func main() {
	stk := new(Stack)
	fmt.Println(stk)
	stk.Push(5)
	fmt.Println(stk)
	stk.Push(1000)
	fmt.Println(stk)
	stk.Push(875)
	fmt.Println(stk)
	stk.Push(7)
	fmt.Println(stk)
	// stk.Push(9) // This returns an error

	p := stk.Pop()
	fmt.Println(p)
	fmt.Println(stk)
	p = stk.Pop()
	fmt.Println(p)
	fmt.Println(stk)
	p = stk.Pop()
	fmt.Println(p)
	fmt.Println(stk)
	p = stk.Pop()
	fmt.Println(p)
	fmt.Println(stk)
	p = stk.Pop() // Returns 0
	fmt.Println(p)
	fmt.Println(stk)
}
