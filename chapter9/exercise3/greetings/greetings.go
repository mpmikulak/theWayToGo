package greetings

import (
	"fmt"
	"time"
)

func hour() int {
	return time.Now().Hour()
}

func Morning() {
	fmt.Println("Good day.")
}

func Night() {
	fmt.Println("Good night.")
}

func IsAfternoon() bool {
	if hour() >= 12 {
		if hour() < 18 {
			return true
		}
		return false
	}
	return false
}

func IsEvening() bool {
	if hour() >= 18 {
		if hour() < 21 {
			return true
		}
		return false
	}
	return false
}

func IsAM() bool {
	if hour() >= 12 {
		return false
	}
	return true
}
