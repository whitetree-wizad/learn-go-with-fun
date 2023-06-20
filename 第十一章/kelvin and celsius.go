package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}
func celsiusToFahrenheit(k float64) float64 {
	return (k * 9.0 / 5.0) + 32.0
}
func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}
func main() {
	var degree float64 = 0.0
	fmt.Printf("kelvin %.2f degree is %.2f degree celsius\n", degree, kelvinToCelsius(degree))
	fmt.Printf(
		"celsius %.2f degree is %.2f degree fahrenheit\n",
		kelvinToCelsius(degree),
		celsiusToFahrenheit(kelvinToCelsius(degree)),
	)
	fmt.Printf("kelvin %.2f degree is %.2f degree Fahrenheit\n", degree, kelvinToFahrenheit(degree))
}
