package main

import (
	"fmt"
	"strconv"
)

const LIMIT = 4

type Stack struct {
	index int
	array [LIMIT]int
}

func newStack() *Stack {
	s := new(Stack)
	s.index = LIMIT - 1
	return s
}

func (s *Stack) Push(n int) {
	s.array[s.index] = n
	s.index--
}

func (s *Stack) Pop() int {
	s.index++
	r := s.array[s.index]
	s.array[s.index] = 0
	return r
}

func (s Stack) String() string {
	str := ""
	for ix, v := range s.array {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(v) + "]"
	}
	return str
}

func main() {
	stk := newStack()
	fmt.Println(stk)
	stk.Push(1)
	fmt.Println(stk)
	stk.Push(2)
	fmt.Println(stk)
	stk.Push(3)
	fmt.Println(stk)
	p := stk.Pop()
	fmt.Println(p)
	fmt.Println(stk)
}
