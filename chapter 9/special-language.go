package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "おはようございます"
	fmt.Println(len(question))
	fmt.Println(utf8.RuneCountInString(question))

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes\n", c, size)

	for i, j := range "おはようございます" {
		fmt.Printf("%v %v %c\n", i, j, j)
	}

}
