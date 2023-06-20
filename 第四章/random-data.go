package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	for count := 0; count != 10; count++ {
		year := rand.Intn(2300) + 1
		mouth := rand.Intn(12) + 1
		dayinmouth := 31
		switch mouth {
		case 2:
			if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
				dayinmouth = 29
			} else {
				dayinmouth = 28
			}
		case 4, 6, 9, 11:
			dayinmouth = 30
		}
		fmt.Println(era, year, mouth, rand.Intn(dayinmouth)+1)
	}
}
