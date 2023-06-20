package main

import "fmt"

func main() {
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%06.2v\n", third)

	ans := 0.1
	var new float64
	for count := 0; count != 11; count++ {
		new += ans
	}
	fmt.Print(new)

}
