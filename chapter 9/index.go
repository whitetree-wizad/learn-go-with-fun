package main

import "fmt"

func main() {
	message := "shalom"
	for count := 0; count != 6; count++ {
		fmt.Printf("%c\n", message[count])
	}
	fmt.Printf("%c", 'g'-'a'+'A')
}
