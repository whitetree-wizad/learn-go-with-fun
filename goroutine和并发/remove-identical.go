package main

import "strings"

func filter_unqual(up, down chan int) {
	his := make(map[int]bool)
	for it := range up {
		if his[it] {
			his[it] = true
			down <- it
		}
	}
	close(down)
}

func split_word(up, down chan string) {
	for it := range up {
		split := strings.Fields(it)
		for _, i := range split {
			down <- i
		}
	}
	close(down)
}
