package main

import (
	"fmt"

	"./stack_general"
)

func main() {
	stk := stack_general.New()

	stk.Push(1)
	stk.Push('a')
	stk.Push("hello, world")
	stk.Push(`s`)

	fmt.Println(stk.Top())
	fmt.Println(stk.Pop())
	fmt.Println(stk.Pop())
	fmt.Println(stk.IsEmpty())
	fmt.Println(stk.Len())
}
