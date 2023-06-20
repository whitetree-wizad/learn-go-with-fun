package main

import (
	"fmt"
	"math/rand"
)

type coordinate struct {
	x, y int
}

type simpmove interface {
	Xmove(int)
	Ymove(int)
}

func move(t simpmove, asix rune, lenth int) {
	switch asix {
	case 'x':
		t.Xmove(lenth)
	case 'y':
		t.Ymove(lenth)
	}
	fmt.Println(t)
}

func (c *coordinate) Xmove(lenth int) {
	c.x += lenth
}

func (c *coordinate) Ymove(lenth int) {
	c.y += lenth
}

func (c *coordinate) String() string {
	return fmt.Sprintf("(%v,%v)", c.x, c.y)
}

func main() {

	turtle := coordinate{x: 2, y: 2}
	for i := 0; i != 10; i++ {
		ranlen := rand.Intn(10)
		ranasix := [2]rune{'x', 'y'}[rand.Intn(2)]

		move(&turtle, ranasix, ranlen)
	}
}
