package main

import (
	"aoc/internal/helpers"
	"bufio"
	"fmt"
	"os"
)

func getInput() [][]int {
	f, _ := os.Open("cmd/day18/input.txt")
	r := bufio.NewReader(f)

	lights := [][]int{}
	for {
		l, _, _ := r.ReadLine()
		if len(l) == 0 {
			break
		}
		line := make([]int, len(l))
		for i, c := range l {
			switch c {
			case '.':
				line[i] = 0
			case '#':
				line[i] = 1
			}
		}
		lights = append(lights, line)
	}
	return lights
}

func getNeighborsOnCount(l [][]int, row, col int) int {
	height := len(l)
	width := len(l[0])

	if 0 < row && row < height-1 && 0 < col && col < width-1 {
		return l[row-1][col-1] + l[row-1][col] + l[row-1][col+1] +
			l[row][col-1] + l[row][col+1] +
			l[row+1][col-1] + l[row+1][col] + l[row+1][col+1]
	} else if row == 0 {
		if col == 0 {
			return l[row][col+1] +
				l[row+1][col] + l[row+1][col+1]
		} else if col == width-1 {
			return l[row][col-1] +
				l[row+1][col-1] + l[row+1][col]
		} else {
			return l[row][col-1] + l[row][col+1] +
				l[row+1][col-1] + l[row+1][col] + l[row+1][col+1]
		}
	} else if row == height-1 {
		if col == 0 {
			return l[row-1][col] + l[row-1][col+1] +
				l[row][col+1]
		} else if col == width-1 {
			return l[row-1][col-1] + l[row-1][col] +
				l[row][col-1]
		} else {
			return l[row-1][col-1] + l[row-1][col] + l[row-1][col+1] +
				l[row][col-1] + l[row][col+1]
		}
	} else if col == 0 {
		return l[row-1][col] + l[row-1][col+1] +
			l[row][col+1] +
			l[row+1][col] + l[row+1][col+1]
	} else {
		return l[row-1][col-1] + l[row-1][col] +
			l[row][col-1] +
			l[row+1][col-1] + l[row+1][col]
	}
}

func nextStep(l [][]int, part2 bool) [][]int {
	copy := make([][]int, len(l))
	for row, line := range l {
		lineCopy := make([]int, len(line))
		for col, light := range line {
			neighborsOn := getNeighborsOnCount(l, row, col)
			if light == 1 {
				if neighborsOn == 2 || neighborsOn == 3 {
					lineCopy[col] = 1
				} else {
					lineCopy[col] = 0
				}
			} else {
				if neighborsOn == 3 {
					lineCopy[col] = 1
				} else {
					lineCopy[col] = 0
				}
			}
		}
		copy[row] = lineCopy
	}

	return copy
}

func Print(l [][]int) {
	for _, line := range l {
		for _, light := range line {
			if light == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func count(l [][]int) int {
	count := 0
	for _, line := range l {
		for _, light := range line {
			count += light
		}
	}
	return count
}

func part1(l [][]int) int {
	for range 100 {
		l = nextStep(l, false)
	}
	return count(l)
}

func updateCorners(l [][]int) [][]int {
	height := len(l)
	width := len(l[0])
	l[0][0] = 1
	l[0][width-1] = 1
	l[height-1][0] = 1
	l[height-1][width-1] = 1
	return l
}

func part2(l [][]int) int {
	updateCorners(l)
	for range 100 {
		l = nextStep(l, false)
		updateCorners(l)
	}
	return count(l)
}

func solution() (int, int) {
	lights := getInput()
	return part1(lights), part2(lights)
}

func main() {
	helpers.PrintResult(solution())
}
