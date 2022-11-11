// Package weather can forecast the weather condition of various cities in Goblinocus.
package weather

// CurrentCondition stores a string for the current condition of the location.
var CurrentCondition string
// CurrentLocation stores a string for the current location. 
var CurrentLocation string

// Forecast ...
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
