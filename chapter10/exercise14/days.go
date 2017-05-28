// Make an alias type Day for int. Define an array of strings with the daynames.
// Define a String() method for type Day which shows the dayname.
// Make an enum const type with iota for all the days of the week (MO, TU, ...)
package main

import "fmt"

var daynames = [7]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

type Day int

const (
	MO Day = iota
	TU
	WED
	TH
	FRI
	SAT
	SUN
)

func (d Day) String() string {
	if d > 6 {
		return ""
	}
	return daynames[d]
}

var daze Day = 9

func main() {
	for i := 0; i < 10; i++ {
		var daze Day = Day(i)
		fmt.Println(daze)
	}
	var day = TU
	fmt.Println(day)
}
