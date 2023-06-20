package main

import (
	"fmt"
)
import "math/rand"

func main() {
	var (
		cent5  = 0.05
		cent10 = 0.1
		cent25 = 0.25
	)
	for ans := 0.0; ans < 20.00; {
		switch rand.Intn(3) {
		case 0:
			ans += cent5
			fmt.Printf("存款%v美元\n", cent5)
		case 1:
			ans += cent10
			fmt.Printf("存款%v美元\n", cent10)
		case 2:
			ans += cent25
			fmt.Printf("存款%v美元\n", cent25)
		}
		fmt.Printf("%0.4f\n", ans)
	}

}
