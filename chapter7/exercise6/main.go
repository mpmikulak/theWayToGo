package main

import (
	"bytes"
	"fmt"
)

func main() {
	n := 5
	var buf bytes.Buffer
	head, tail := buf.Bytes()[0:n], buf.Bytes()[n:len(buf.Bytes())]

	fmt.Println(tail)
	fmt.Println(head)
}
