package main

import "fmt"

const (
	HOUR TZ = 1
	UTC  TZ = HOUR * 0
	EST  TZ = HOUR * -4
	CST  TZ = HOUR * -5
	MST  TZ = HOUR * -6
	PST  TZ = HOUR * -7
)

type TZ int

var timezones = map[TZ]string{
	UTC: "Universal Greenwhich Time",
	EST: "Eastern Standard Time",
	CST: "Central Standard Time",
	MST: "Mountain Standard Time",
	PST: "Pacific Standard Time",
}

func (tz TZ) String() string {
	for name, zone := range timezones {
		if tz == name {
			return zone
		}
	}
	return ""
}

func main() {
	fmt.Println(HOUR * -4)
}
