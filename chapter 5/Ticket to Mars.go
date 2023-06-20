package main

import (
	"fmt"
	"math/rand"
)

var distance = 62100000

func main() {
	fmt.Printf("%-20v %-10v %-10v %-10v\n", "airline", "duration", "ticket", "price")

	for count := 0; count != 10; count++ {
		var airline = ""
		switch rand.Intn(3) {
		case 0:
			airline = "Space Adventures"
		case 1:
			airline = "SpaceX"
		case 2:
			airline = "Virgin Galactic"
		}

		var speeds = rand.Intn(15) + 16
		var flydur = distance / speeds / 3600 / 24
		var price = speeds + 20 + rand.Intn(2)

		var flysel = ""
		switch rand.Intn(2) {
		case 0:
			flysel = "单程"
		case 1:
			flysel = "往返"
			price *= 2
		}

		fmt.Printf("%-20v %-10v %-10v %-10v\n", airline, flydur, flysel, price)
	}
}
