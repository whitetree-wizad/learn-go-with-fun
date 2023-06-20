package main

import "fmt"

func main() {

	var age = 21
	var minor = age < 18

	fmt.Printf("At age %v, am I a minor? %v", age, minor)

}
