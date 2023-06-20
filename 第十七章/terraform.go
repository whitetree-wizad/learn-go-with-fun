package main

import "fmt"

type StringSlice []string

func add_new(s string) string {
	return "New" + "_" + s
}

func (p StringSlice) terraform() {
	for i, value := range p {
		p[i] = add_new(value)
	}
}

func main() {
	plants := []string{"Mars", "Uranus", "Neptune"}
	StringSlice(plants).terraform()
	fmt.Println(plants)
}
