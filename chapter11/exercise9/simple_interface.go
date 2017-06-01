package main

import "fmt"

// Simpler interface and methods
type Simpler interface {
	Get() int
	Set(int)
}

func fI(s Simpler) int {
	switch s.(type) {
	case *Simple:
		s.Set(6)
		return s.Get()
	case *RSimple:
		s.Set(7)
		return s.Get()
	default:
		return 0
	}
}

// Empty interface method 
func gI(i interface{}) int{
	if v, ok := i.(Simpler); ok{
		return v.Get()
	}
	return 0
}

// RSimple definition and methods
type RSimple struct {
	number int
}

func (r *RSimple) Get() int {
	return r.number
}

func (r *RSimple) Set(i int) {
	r.number = i
}

// Simple Definition and methods
type Simple struct {
	number int
}

func (s *Simple) Get() int {
	return s.number
}

func (s *Simple) Set(i int) {
	s.number = i
}

func main() {
	first := new(Simple)
	second := new(RSimple)

	first.Set(8)
	fmt.Println(gI(first))
	fmt.Println(gI(second))
}
