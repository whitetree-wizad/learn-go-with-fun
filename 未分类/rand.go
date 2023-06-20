package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num)

	var mindistance, randadvance = 56000000, 401000000 - 56000000 + 1 //km
	fmt.Println(mindistance + rand.Intn(randadvance))

	fmt.Println(56000000 / 28 / 24)
}
