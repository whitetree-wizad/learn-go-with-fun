package main

import (
	"math/rand"
)

func stillapp(slice []int) {
	capacity := cap(slice)
	for capacity <= 1000 {
		if capacity >= cap(slice) {
			slice = append(slice, rand.Int())
		} else {
			capacity = cap(slice)
			println(capacity, slice)
		}
	}
}

func main() {
	new_slice := make([]int, 3, 3)
	stillapp(new_slice)
}
