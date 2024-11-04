package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
)

func getInput() <-chan string {
	ch := make(chan string)

	go func() {
		f, _ := os.Open("cmd/day8/input.txt")
		r := bufio.NewReader(f)

		for {
			l, _, _ := r.ReadLine()
			if len(l) == 0 {
				break
			}
			ch <- string(l)
		}

		close(ch)
	}()

	return ch
}

type Line struct {
	Raw         string
	Len         int
	MemoryPart1 int
	MemoryPart2 int
}

func (l Line) DifPart1() int {
	return l.Len - l.MemoryPart1
}

func (l Line) DifPart2() int {
	return l.MemoryPart2 - l.Len
}

func solution() (int, int) {
	part1 := 0
	part2 := 0

	for l := range getInput() {
		line := Line{l, len(l), 0, 2}
		for i := 0; i < len(l); i++ {
			c := l[i]
			switch c {
			case '\\':
				line.MemoryPart2 += 2
				switch l[i+1] {
				case '"':
					line.MemoryPart1++
					line.MemoryPart2 += 2
					i++
				case 'x':
					line.MemoryPart1++
					line.MemoryPart2 += 3
					i += 3
				case '\\':
					line.MemoryPart1++
					line.MemoryPart2 += 2
					i++
				}
			case '"':
				line.MemoryPart2 += 2
			default:
				line.MemoryPart1++
				line.MemoryPart2++
			}
		}
		part1 += line.DifPart1()
		part2 += line.DifPart2()
	}

	return part1, part2
}

func main() {
	helpers.PrintResult(solution())
}
