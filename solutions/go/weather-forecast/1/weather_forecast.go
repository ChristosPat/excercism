//Package weatheris a program that tell us the weather.
package weather

//CurrentCondition gives the weather now.
var CurrentCondition string
//CurrentLocation tell us the location that weather is about.
var CurrentLocation string


/* Forecast return a string with the
Location-city condition and weather conditoin.*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
