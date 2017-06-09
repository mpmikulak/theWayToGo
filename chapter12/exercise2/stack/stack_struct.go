package stack

import "fmt"

type Stack struct {
	Index int
	Value []int
}

func New() *Stack {
	s := new(Stack)
	for i := 0; i < 3; i++ {
		s.Value = append(s.Value, 0)
	}
	return s
}

// Push moves all elements of the slice up one spot and places the input in the
// zero index
func (s *Stack) Push(n int) {
	if s.Index == len(s.Value) {
		s.Value = append(s.Value, 0)
	}
	for i := s.Index; i > 0; i-- {
		s.Value[i] = s.Value[i-1]
	}
	s.Value[0] = n
	s.Index++
}

// Pull moves all elements from index 2 and above down one index
func (s *Stack) Pull() {
	if s.Index >= 1 {
		for i := 1; i < s.Index; i++ {
			s.Value[i-1] = s.Value[i]
		}
		s.Index--
		return
	}
	s.Value[0] = 0
	s.Index--
}

// Delete zeroes the top-most element and does not return or affect the index
func (s *Stack) Delete() {
	s.Value[s.Index] = 0
}

// Pop removes the top-most element and returns it. Also deducts from the index
func (s *Stack) Pop() (n int) {
	n = s.Value[s.Index]
	s.Value[s.Index] = 0
	s.Index--
	return
}

// String is used for prettifying output and troubleshooting
func (s *Stack) String() string {
	str := fmt.Sprintf("%v: %v", s.Index, s.Value)
	return str
}
