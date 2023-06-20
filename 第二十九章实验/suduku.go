package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const row, col = 9, 9

type intmap map[int]bool
type suduku [row][col]int
type suduerr []error

var (
	inbounderr  = errors.New("entered number out of bounds")
	inroworcol  = errors.New("rows or cols repeated")
	cantediterr = errors.New("can not modify this digit")
)

type suduku99 struct {
	sudu suduku
	err  suduerr
}

func (s suduku) read(rows, cols int) int {
	abs := s[rows][cols]
	if abs < 0 {
		return -abs
	}
	return abs
}

func gersuduku() suduku {
	var s suduku
	for i := range s {
		for j := range s[i] {
			s[i][j] = 0
		}
	}
	return s
}

func (s *suduku) set(digit, rows, cols int) suduerr {
	var err suduerr
	if !s.canedit(rows, cols) {
		err = append(err, cantediterr)
	}
	if !inBound(rows, cols) {
		err = append(err, inbounderr)
	}
	if !inrowsum(digit, cols, *s) || !incolsum(digit, rows, *s) {
		err = append(err, inroworcol)
	}

	if len(err) != 0 {
		return err
	}
	s[rows][cols] = -digit
	return nil
}

func (s *suduku) clear(rows, cols int) suduerr {
	var err suduerr
	s[rows][cols] = 0
	if !inBound(rows, cols) {
		err = append(err, inbounderr)
	}

	if len(err) != 0 {
		return err
	}
	return nil
}

func (s suduku) canedit(rows, cols int) bool {
	if s[rows][cols] > 0 {
		return false
	}
	return true
}

func inBound(rows, cols int) bool {
	if rows < 0 || rows > row || cols < 0 || cols > col {
		return false
	}
	return true
}

func inrowsum(new, cols int, s suduku) bool {
	sum := make(intmap)
	for j, _ := range s[cols] {
		sum[s.read(j, new)] = true
	}
	old := len(sum)
	sum[new] = true
	if !(len(sum) > old) {
		return false
	}
	return true
}

func incolsum(new, rows int, s suduku) bool {
	sum := make(intmap)
	for j, _ := range s[rows] {
		sum[s.read(new, j)] = true
	}
	old := len(sum)
	sum[new] = true
	if !(len(sum) > old) {
		return false
	}
	return true
}

func (err suduerr) Error() string {
	var s []string
	for _, err := range err {
		s = append(s, err.Error())
	}
	return strings.Join(s, ",")
}
func (s suduku99) String() string {
	top := fmt.Sprintf("suduku:\n")
	for rows, i := range s.sudu {
		for cols := range i {
			top += fmt.Sprintf("%v\t", s.sudu.read(rows, cols))
		}
		top += "\n"
	}
	return top
}

func main() {
	newsuduku := &suduku99{gersuduku(), nil}
	newsuduku.err = newsuduku.sudu.set(1, 1, 1)
	newsuduku.err = newsuduku.sudu.set(7, 3, 8)
	if newsuduku.err != nil {
		fmt.Printf("%v", newsuduku.err.Error())
		os.Exit(1)
	}
	fmt.Println(newsuduku)
}
