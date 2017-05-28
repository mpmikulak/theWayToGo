package main

import "container/list"

func (p *list.List) Iter() { // Cannot define a method on a type thats not in this package
	//
}

func main() {
	lst := new(list.List)
	for _ = range lst.Iter() {

	}
}
