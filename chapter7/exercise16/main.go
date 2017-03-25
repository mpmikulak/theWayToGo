// Traverses an array of characters and copies to another array only if
// the character is different from the one that precedes it.
package main

import "fmt"

func main() {
	str := "hello? are all leet peeps at the pool?"
	str2 := unique([]byte(str))
	fmt.Println(string(str2))

}

func unique(s []byte) []byte {
	retStr := make([]byte, 0)
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			retStr = append(retStr, s[i])
		}
	}
	return retStr
}
