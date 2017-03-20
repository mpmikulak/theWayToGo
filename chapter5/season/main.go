// Reports what season it is based on system time
package main

import (
	"fmt"
	"time"
)

func season() string {
	month := time.Now().Month()
	switch month {
	case time.January, time.February, time.December:
		return "It's Winter"
	case time.March, time.April, time.May:
		return "It's Spring"
	case time.June, time.July, time.August:
		return "It's Summer"
	case time.September, time.October, time.November:
		return "It's Fall"
	}
	return "OH NOES!"
}

func main() {
	fmt.Println(season())
}
