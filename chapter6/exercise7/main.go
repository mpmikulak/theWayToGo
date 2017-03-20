// Makes all non ASCII characters a '?'
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, 世界"
	result := strings.Map(questionMark, str)
	fmt.Println(result)

}

func questionMark(k rune) rune {
	if k > 127 {
		return '?'
	}
	return k
}
