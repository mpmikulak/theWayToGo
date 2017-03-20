// Calculates the square root of an input and returns an error if negative
package main

import (
	"errors"
	"fmt"
	"math"
)

var first, second float64 = 4.0, -800.0

func main() {
	num, err := MySqrt(first)
	if err != nil {
		fmt.Println(err, num)
	} else {
		fmt.Println(num)
	}

	num, err = MySqrt(second)
	if err != nil {
		fmt.Println(err, num)
	} else {
		fmt.Println(num)
	}

}

func MySqrt(x float64) (y float64, err error) {
	if x < 0 {
		err = errors.New("Negative number in square root function.")
		return 1.0, err
	}
	return math.Sqrt(x), nil
}
