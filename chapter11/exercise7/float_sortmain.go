package main

import (
	"fmt"

	"./float64"
)

func main() {
	float := float64.New()
	float.Fill()
	float64.Sort(float)
	fmt.Println(float)
}
