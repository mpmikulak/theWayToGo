// Make a map to hold together the number of the day in the week with its name
package main

import "fmt"

func main() {
	days := map[string]int{"Monday": 1, "Tuesday": 2, "Wednesday": 3,
		"Thursday": 4, "Friday": 5, "Saturday": 6, "Sunday": 7}

	for key, value := range days {
		fmt.Printf("%v is day number %v.\n", key, value)
	}
	if day, ok := days["tuesday"]; ok {
		fmt.Println(day)
	}

	if day, ok := days["hollyday"]; ok {
		fmt.Println(day)
	} else {
		fmt.Println("That day does not exist.")
	}
}
