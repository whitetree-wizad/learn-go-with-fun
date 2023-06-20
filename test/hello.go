package main

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}
func main() {
	terry := &person{
		name: "Terry",
		age:  15,
	}

}
