package main

import (
	"fmt"
	"time"

	"./greetings"
)

func main() {

	if time.Now().Hour() >= 12 {
		greetings.Night()
	} else {
		greetings.Morning()
	}

	fmt.Println(greetings.IsAM())
	fmt.Println(greetings.IsAfternoon())
	fmt.Println(greetings.IsEvening())
}
