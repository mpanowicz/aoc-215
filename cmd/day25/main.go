package main

import (
	"aoc/internal/helpers"
)

const (
	multiply = 252533
	modulo   = 33554393
)

func solution() (int, int) {
	row := 2981
	col := 3075

	firstRowColumnNumber := (1 + col) * col / 2
	codeNumber := firstRowColumnNumber + (row-1)*col + (1+row-2)*(row-2)/2

	start := 20151125
	for i := 2; i <= codeNumber; i++ {
		a := start % modulo
		b := multiply % modulo
		c := a * b
		start = c % modulo
	}

	return start, 0
}

func main() {
	helpers.PrintResult(solution())
}
