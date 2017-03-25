// If str is a string, what then is str[len(str)/2:] + str[:len(str)/2]?
// Test it !
package main

import "fmt"

func main() {
	str := "What is love?"
	newStr := str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(str)
	fmt.Println(newStr)
}
