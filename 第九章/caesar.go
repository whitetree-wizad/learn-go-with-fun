package main

import "fmt"

func main() {
	question := "L fdph,L vdz,L frqtxhuhg"

	for i := 0; i != len(question); i++ {
		w := question[i]
		if w >= 'a' && w <= 'z' {
			w -= 3
			if w < 'a' {
				w += 26
			}
		}
		if w >= 'A' && w <= 'Z' {
			w -= 3
		}
		switch {
		case w == 'a':
			w = 'x'
		case w == 'b':
			w = 'y'
		case w == 'c':
			w = 'z'
		}
		fmt.Printf("%c", w)
	}
}
