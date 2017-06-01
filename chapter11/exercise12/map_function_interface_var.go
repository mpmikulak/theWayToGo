package main

import "fmt"

type obj interface{}

func mapFunc(f func(obj) obj, o ...[]obj) {
	for _, out := range o {
		for ix, v := range out {
			out[ix] = f(v)
		}
	}
}

func main() {
	double := func(o obj) obj {
		switch o.(type) {
		case int:
			return o.(int) * 2
		case string:
			return o.(string) + o.(string)
		}
		return o
	}

	strings := []obj{"alalal", "fhfhfh", "llaaallskdh", "qwertysddf", 45, 65, 78, 89, 45, 1}

	mapFunc(double, strings)
	fmt.Println(strings)
}
