// Stack is used to initialize and manipulate a data structure
package stack

import "strconv"

const LIMIT = 4

// Stack is a data structure that holds an array and an index of the next empty
// cell
type Stack struct {
	index int
	array [LIMIT]int
}

// New returns a newly initialized pointer to a stack
func New() *Stack {
	s := new(Stack)
	s.index = LIMIT - 1
	return s
}

// Push finds the highest empty cell and fills it with the argument
func (s *Stack) Push(n int) {
	s.array[s.index] = n
	s.index--
}

// Pop returns the value in the lowest filled cell and then clears it
func (s *Stack) Pop() int {
	s.index++
	r := s.array[s.index]
	s.array[s.index] = 0
	return r
}

// String is used for debugging purposes
func (s Stack) String() string {
	str := ""
	for ix, v := range s.array {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(v) + "]"
	}
	return str
}
