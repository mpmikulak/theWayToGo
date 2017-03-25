// defines a function that removes items from start index to finish index
package main

import "fmt"

func main() {
	str := []string{"h", "e", "l", "l", "o", "w", "o", "r", "l", "d"}
	fmt.Printf("%s\n", str)
	str = RemoveStringSlice(str, 2, 4)
	fmt.Printf("%s\n", str)

}

func RemoveStringSlice(str []string, start int, end int) []string {
	newStr := make([]string, 0)
	for i := range str {
		if i >= start && i <= end {
			continue
		}
		newStr = append(newStr, str[i])
	}
	return newStr
}
