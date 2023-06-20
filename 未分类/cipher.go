package main

import "fmt"

func main() {
	plainText := "your message goes here"
	key := "GOLANG"

	var t int32 = 0
	var word int32 = 0
	for _, c := range plainText {
		if c >= 'a' && c <= 'z' {
			t %= 6
			word = int32(key[t])
			t++

			c = (c+word-'A'-'a')%26 + 'a'
		}
		fmt.Printf("%c", c)
	}
}
