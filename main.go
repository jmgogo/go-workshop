package main

import "type_definitions/types"

func main() {
	println("Trying custom types with methods:")
	types.MilesToKilometers(types.DistanceMiles(10))
	types.KilometersToMiles(types.DistanceKm(10))

	// initialize john and upgrade him as an instructor
	john := types.NewUser("John", "john@example.com", 30)
	instructorJohn := types.NewInstructor(john, "Math", "Room 101")

	// initialize billy as basic user
	billy := types.User{Name: "Billy", Email: "bill@example.com", Age: 25}

	var people [2]types.PrettyPrintable

	people[0] = billy
	people[1] = instructorJohn

	for i, person := range people {
		println("\nPerson ", i)
		person.PrettyPrint()
	}
}
