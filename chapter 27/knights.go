package main

import "fmt"

type character struct {
	name     string
	leftHand item
}

type item string

func (c character) String() string {
	var none item
	if c.leftHand == none {
		return fmt.Sprintf("%v takes nothing\n", c.name)
	}
	return fmt.Sprintf("%v takes %v\n", c.name, c.leftHand)
}

func (c *character) pickup(i *item) {
	c.leftHand = *i
}

func (c *character) give(to *character) {
	var none item
	to.leftHand = c.leftHand
	c.leftHand = none
}

func main() {
	var none item
	aser := character{name: "aser", leftHand: none}
	knight := character{name: "knights", leftHand: none}

	pre := item("Caliburnus")
	knight.pickup(&pre)
	knight.give(&aser)

	fmt.Println(aser, knight)
}
