package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const rows, columns = 9, 9

var (
	errbounds = errors.New("out of bounds")
	errdigit  = errors.New("invalid digit")
)

type Grid [rows][columns]int8

type sudokuError []error

func (se sudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ",")
}

func (g *Grid) set(row, column int, digit int8) error {
	var errs sudokuError
	if !inBounds(row, column) {
		errs = append(errs, errbounds)
	}
	if !invalid(row, column) {
		errs = append(errs, errdigit)
	}

	if len(errs) > 0 {
		return errs
	}
	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column > columns {
		return false
	}
	return true
}

func invalid(row, colum int) bool {
	if row < 1 || row > 9 || colum < 1 || colum > 9 {
		return false
	}
	return true
}

func main() {
	var g *Grid
	g = new(Grid)
	err := g.set(10, 4, 5)
	if err != nil {
		if errs, ok := err.(sudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
	}
	os.Exit(1)
}
