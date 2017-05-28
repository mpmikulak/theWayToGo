package main

import (
	"fmt"
	"strconv"
)

type T struct {
	a int
	b float64
	c string
}

func (t *T) String() string {
	intString := strconv.Itoa(t.a)                      // Converts the int to a string
	floatString := strconv.FormatFloat(t.b, 'f', 6, 64) // Converts the float to a string with specified precision
	quoteString := strconv.Quote(t.c)                   // Converts the interpreted string to a string literal

	return intString + " / " + floatString + " / " + quoteString
}

func main() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Println(t.String())
}
