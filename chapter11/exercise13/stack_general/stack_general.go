package stack_general

import (
	"errors"
)

type obj interface{}

type Stack struct {
	index int
	data  []obj
}

func New() *Stack {
	return new(Stack)
}

func (s Stack) Len() int {
	return s.index
}

func (s Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(x obj) {
	s.data = append(s.data, x)
	s.index++
}

func (s *Stack) Pop() (x obj, e error) {
	if s.index == 0 {
		e = errors.New("This array is full!")
		x = 0
		return
	}
	s.index--
	x = s.data[s.index]
	e = nil
	s.data[s.index] = 0
	return
}

func (s Stack) Top() (x obj, e error) {
	if s.index == 0 {
		e = errors.New("This array is full!")
		x = 0
		return
	}
	x = s.data[s.index-1]
	e = nil
	return
}
