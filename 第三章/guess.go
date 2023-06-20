package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		var (
			randnum  = rand.Intn(100) + 1
			guessnum = 21
		)
		if randnum > guessnum {
			fmt.Println(randnum, "more.")
		} else if randnum < guessnum {
			fmt.Println(randnum, "less.")
		} else {
			fmt.Println("bingo!")
			break
		}
	}

}
