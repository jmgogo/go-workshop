package birthday

import "time"

func Generation(age *int) {
	// get the current year
	currentYear := time.Now().Year()

	// calculate the birth year
	birthYear := currentYear - *age

	// determine the generation based on the birth year
	switch {
	case birthYear >= 1997 && birthYear <= 2012:
		println("You belong to Generation Z.")
	case birthYear >= 1981 && birthYear <= 1996:
		println("You belong to the Millennial Generation.")
	case birthYear >= 1965 && birthYear <= 1980:
		println("You belong to Generation X.")
	case birthYear >= 1946 && birthYear <= 1964:
		println("You belong to the Baby Boomer Generation.")
	default:
		println("Your generation is not categorized.")
	}

}
