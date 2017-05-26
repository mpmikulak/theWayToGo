// Construct a collection which maps English names of drinks to the French
// (or your native language) translations; print first only the drinks
// available, and then print both (the name and the translation). Then
// produce the same output, but this time the English names of the drinks
// must be sorted.
package main

import (
	"fmt"
	"sort"
)

var drinks = map[string]string{"Tea": "Thé", "Wine": "Le vin",
	"Beer": "Bièr", "Coffee": "Le Café", "Apple Juice": "Le jus de pomme",
	"Milk": "Le lait", "Mineral Water": "l'eau minéral"}

func main() {
	for key := range drinks {
		fmt.Println(key)
	}

	for key := range drinks {
		fmt.Printf("%v is %v in french.\n", key, drinks[key])
	}
	fmt.Println()

	i := 0
	keys := make([]string, len(drinks))
	for key := range drinks {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%v is %v in French\n", k, drinks[k])
	}
}
