package main

import "fmt"

func main() {
	message := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i != len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		/*		switch {
				case c >= 'n':
					c -= 13
				case c >= 'a' && c < 'n':
					c += 13
				}
		*/
		fmt.Printf("%c", c)
	}

}
