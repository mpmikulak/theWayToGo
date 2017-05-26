package main

import (
	"fmt"
	"unsafe"
)

func main() {
	number := 40000

	size := unsafe.Sizeof(number)
	fmt.Println(size)
}
