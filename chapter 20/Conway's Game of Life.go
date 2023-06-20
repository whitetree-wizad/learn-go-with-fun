package main

import (
	"fmt"
	"math/rand"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool
type UniShow [][]string
type coordinate struct {
	x, y int
}

func newCoordinate(x, y int) coordinate {
	return coordinate{x, y}
}

func transCoordinate(co coordinate) coordinate {
	return coordinate{co.y - 1, co.x - 1}
}

func main() {
	Uni := NewUniverse()
	Uni.Seed(77)
	Uni.Show()
	testpoint := newCoordinate(3, 1)
	fmt.Println(Uni.alive(testpoint))
	Uni.Neighbors(testpoint)
}

func NewUniverse() Universe {
	NewUni := make(Universe, width)
	for i := range NewUni {
		NewUni[i] = make([]bool, height)
	}
	for i := 0; i != width; i++ {
		for j := 0; j != height; j++ {
			NewUni[i][j] = false
		}
	}
	return NewUni
}

func (u Universe) Show() {
	for _, q := range u {
		for _, p := range q {
			if p {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Print("\n")
	}
}

func (u Universe) Seed(n int64) {
	rand.Seed(n)
	for _, q := range u {
		for i, _ := range q {
			ran := rand.Intn(100) <= 25
			if ran {
				q[i] = true
			} else {
				q[i] = false
			}
		}
	}
}

func (u Universe) alive(co coordinate) bool {
	co.x = (co.x + 15) % 15
	co.y = (co.y + 80) % 80
	transco := transCoordinate(co)
	return u[transco.x][transco.y]
}

func (u Universe) Neighbors(co coordinate) int {
	var sum int = 0
	for _, i := range [3]int{co.x - 1, co.x, co.x + 1} {
		for _, j := range [3]int{co.y - 1, co.y, co.y + 1} {
			sideco := newCoordinate(i, j)
			if u.alive(sideco) {
				sum++
			}
		}
	}
	return sum - 1
}
