package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
