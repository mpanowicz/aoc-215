package main

import (
	"aoc/internal/helpers"
	"bufio"
	"math"
	"os"
	"sort"
)

func getInput() []int {
	f, _ := os.Open("cmd/day17/input.txt")
	r := bufio.NewReader(f)

	sizes := []int{}
	for {
		l, _, _ := r.ReadLine()
		if len(l) == 0 {
			break
		}

		sizes = append(sizes, helpers.ParseInt(string(l)))
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return sizes
}

func fillContainers(sizes []int, remaining int) [][]int {
	solutions := [][]int{}
	for i := 0; i < len(sizes); i++ {
		if sizes[i] == remaining {
			solutions = append(solutions, []int{sizes[i]})
		} else if sizes[i] < remaining {
			partRemaining := remaining - sizes[i]
			rest := fillContainers(sizes[i+1:], partRemaining)
			for _, c := range rest {
				if helpers.SumInt(c) == partRemaining {
					solutions = append(solutions, append([]int{sizes[i]}, c...))
				}
			}
		}
	}
	return solutions
}

func solution() (int, int) {
	sizes := getInput()
	containers := fillContainers(sizes, 150)
	part1 := len(containers)

	count := 0
	min := math.MaxInt
	for _, c := range containers {
		l := len(c)
		if l < min {
			min = l
			count = 1
		} else if l == min {
			count++
		}
	}
	part2 := count

	return part1, part2
}

func main() {
	helpers.PrintResult(solution())
}
