package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type RSimple struct {
	number int
}

func (r *RSimple) Get() int {
	return r.number
}

func (r *RSimple) Set(i int) {
	r.number = i
}

type Simple struct {
	number int
}

func (s *Simple) Get() int {
	return s.number
}

func (s *Simple) Set(i int) {
	s.number = i
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

func main() {
	first := new(Simple)
	second := new(RSimple)
	fmt.Println(fI(first))
	fmt.Println(fI(second))
}
