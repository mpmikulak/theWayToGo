// Use the package container/list to construct a double linked list, put
// the values 101,102,103 in it and make a printout of the list.
package main

import (
	"container/list"
	"fmt"
)

func main() {
	numbers := []int{101, 102, 103}
	dlist := list.New()
	fmt.Println(dlist)

	for _, i := range numbers {
		fmt.Println(i)
		dlist.PushBack(i)
	}
	fmt.Println(dlist)

	for e := dlist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
