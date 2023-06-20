package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)

		var fail = rand.Intn(100)
		if fail == 0 {
			fmt.Println("lift failed.")
			break
		}

		count--
	}
	if count == 0 {
		fmt.Println("liftoff!")
	}
}
