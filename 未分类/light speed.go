package main

import "fmt"

func main() {
	const Lightspeed = 299792 //km/s
	var distance = 56000000   //km
	fmt.Println(distance/Lightspeed, "seconds")
	distance = 401000000
	fmt.Println(distance/Lightspeed, "seconds")

	distance = 96300000
	var Rocketspeed = 100800 / 24 //km/d
	fmt.Println(distance/Rocketspeed, "days")

}
