package birthday

import "fmt"

func CheckValidDriversLicense(age int) {
	if age >= 100 {
		panic("You cannot drive at this age!")
	}

	if age < 16 {
		fmt.Println("No driving for you yet!")
	} else {
		fmt.Println("Valid driver's license.")
	}
}
