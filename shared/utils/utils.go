package utils

import "math"

// Round a floating point number
func Round(f float64) float64 {
	return math.Floor(f + .5)
}

// Round a floating point number to a specific number of places
func RoundToPlaces(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return Round(f * shift) / shift
}
