// Make an alias type Celsius for float64 and define a String() method for it
// which prints out the temperature with 1 decimal and °C.
package main

import "fmt"

type Celsius float64

func (c *Celsius) String() {
	fmt.Printf("%.1f\u00BAC\n", *c)
}

var degree Celsius = 5.223

func main() {
	degree.String()
}
