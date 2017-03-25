// Defines a functio Split which splits a string at the given position
// and returns the two sub strings
package main

import "fmt"

func main() {
	str := "My milkshake is better"
	fmt.Println(str)
	s1, s2 := Split(str, -8)
	fmt.Println(s1)
	fmt.Println(s2)

}

func Split(s string, p int) (string, string) {
	if p <= 0 {
		return "", ""
	}

	str1 := make([]byte, p)
	str2 := make([]byte, len(s)-p)
	b := []byte(s)

	str1 = b[:p]
	str2 = b[p:]
	return string(str1), string(str2)
}
