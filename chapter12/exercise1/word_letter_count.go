// Word letter count reads text from the keyboard and counts the number of
// characters including spaces, the number of words, and the number of lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	input, err := read.ReadString('S')
	if err != nil {
		return
	}

	lines, words, char := count(input)

	fmt.Printf("Lines : %v\n", lines)
	fmt.Printf("Words: %v\n", words)
	fmt.Printf("Characters: %v\n", char)
}

func count(s string) (lines, words, char int) {
	for _, v := range s {
		switch {
		case v == '\n':
			lines++
			words++
		case v == ' ':
			words++
			char++
		default:
			char++
		}
	}
	char--
	return
}
