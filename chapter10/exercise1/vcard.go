// Define a struct Address and a struct VCard. The latter contains a
// person’s name, a number of addresses, a birth date, a photo. Try to find
// the right data types. Make your own vcard and print its contents.
package main

import "fmt"

// Address holds all of the values of a person's address
type Address struct {
	number int
	street string
	city   string
	state  string
	zip    int
}

// VCard holds the values for a person's VCard
type VCard struct {
	name     string
	address  *Address
	birthday string
	photo    string
}

func main() {
	myAddress := &Address{305, "Janie", "Smyrna", "Tennessee", 37167}
	myName := "Mike Mikulak"
	myBirthday := "09/30/90"
	myPhoto := "yeah, right"

	me := VCard{myName, myAddress, myBirthday, myPhoto}

	fmt.Println(me)
}
