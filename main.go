package main

import (
	"control_flows/birthday"
	"fmt"
)

var age int = 15

func main() {

	birthday.Generation(&age)
	birthday.CheckValidDriversLicense(age)

	// this is deferred to run at the end of the function
	defer fmt.Printf("You were %d years old.\n", age)

	yearsAroundSun := 5
	fmt.Printf("Let's go %d years around the sun.\n", yearsAroundSun)

	// many years around the sun... ☀️
	for i := 0; i < yearsAroundSun; i++ {
		birthday.Birthday(&age)
	}

	fmt.Printf("You are now %d years old.\n", age)
	birthday.CheckValidDriversLicense(age)
}
