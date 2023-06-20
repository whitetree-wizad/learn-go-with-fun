package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func realsenser() kelvin {
	return 0
}

func fakesensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var temp kelvin = 5
	sensor := calibrate(realsenser, temp)
	temp++
	fmt.Println(sensor())

	fsensor := calibrate(fakesensor, temp)
	fmt.Println(fsensor(), fsensor(), fsensor())

}
