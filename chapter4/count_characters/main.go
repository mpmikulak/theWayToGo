package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "asSASA ddd dsjkdsjs dk"
	str2 := "asSASA ddd dsjkdsjsこん dk"

	lenStr1 := len(str1)
	lenStr2 := len(str2)
	fmt.Println("Number of bytes in str1: ", lenStr1)
	fmt.Println("Number of bytes in str2: ", lenStr2)

	runes1 := utf8.RuneCountInString(str1)
	runes2 := utf8.RuneCountInString(str2)
	fmt.Println("Number of runes in str1", runes1)
	fmt.Println("Number of runes in str2", runes2)
}
