package types

// Example of type aliasing
type json = map[string]string

// Creating a new type that offers semantic meaning
type DistanceMiles float64 // miles
type DistanceKm float64    // kilometers

// with a custom type you can add methods to it
func (d DistanceMiles) toKilometers() DistanceKm {
	return DistanceKm(d * 1.60934)
}

func (d DistanceKm) toMiles() DistanceMiles {
	return DistanceMiles(d / 1.60934)
}

func MilesToKilometers(d DistanceMiles) {
	println("Distance in kilometers: ", d.toKilometers())
}

func KilometersToMiles(d DistanceKm) {
	println("Distance in miles: ", d.toMiles())
}
