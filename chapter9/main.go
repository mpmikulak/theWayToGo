package main

import (
	"fmt"
	"regexp"
	"strconv"

	"./pack1"
)

func main() {
	// REGEXP
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18" // String to search in
	pat := "[0-9]+.[0-9]+"                                      // Pattern to match

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok { // Test for a match
		fmt.Println("Match found!!")
	}

	re, _ := regexp.Compile(pat) // Compile the pattern into a *Regexp

	str := re.ReplaceAllString(searchIn, "##.#") // Replace the pattern with "##.#"
	fmt.Println(str)

	// Using a function
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)

	// Importing custom made packages
	integer := pack1.Pack1Int
	fmt.Println(integer)
}
