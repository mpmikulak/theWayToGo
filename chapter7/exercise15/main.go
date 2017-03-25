// Reverses a string
package main

import "fmt"

func main() {
	str := "I want to know what love is"
	newStr := reverse(str)
	fmt.Println(newStr)

}

func reverse(s string) string {
	b := []byte(s)
	for i := range b {
		b[i], b[(len(b)-i)-1] = s[(len(b)-i)-1], s[i]
	}
	return string(b)
}
