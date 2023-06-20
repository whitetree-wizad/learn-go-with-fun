package main

import (
	"fmt"
	"math/rand"
)

type Stringer interface {
	String() string
}

type motivate interface {
	move()
	eat()
}

type animal struct {
	name, favorfood string
}

func (a animal) String() string {
	return a.name
}

func (a animal) move() {
	rnum := rand.Intn(3)
	switch rnum {
	case 0:
		fmt.Printf("%v is playing tennis.", a.name)
	case 1:
		fmt.Printf("a lot of bees are running after %v", a.name)
	case 2:
		fmt.Printf("%v is taking shower", a.name)
	}
	fmt.Printf("\n")
}

func (a animal) eat() {
	fmt.Printf("%v is eating its favorate food:%v\n", a.name, a.favorfood)
}

func main() {
	rand.Seed(32)
	dogs := animal{"dog", "meat"}
	cats := animal{"cat", "fish"}
	goats := animal{"goat", "grass"}

	for hour := 0; hour != 72; hour++ {
		rnum := rand.Intn(3)
		var ani animal
		switch rnum {
		case 0:
			ani = dogs
		case 1:
			ani = cats
		case 2:
			ani = goats
		}

		var t motivate
		t = ani

		rnum = rand.Intn(2)

		switch rnum {
		case 0:
			t.move()
		case 1:
			t.eat()
		}
	}
}
