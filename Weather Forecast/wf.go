// Package weather created to give forecast conditions.
package weather

// CurrentCondition rapresents the current condition as a string.
var CurrentCondition string

// CurrentLocation rapresent the current location as a string.
var CurrentLocation string

// Forecast returns a string containing the location and the current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
