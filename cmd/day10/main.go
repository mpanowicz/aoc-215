package main

import (
	"aoc/internal/helpers"
	"fmt"
	"strings"
)

func process(s string) string {
	var sb strings.Builder
	c := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		if c == s[i] {
			count++
		} else {
			sb.WriteString(fmt.Sprintf("%d%s", count, string(c)))
			c = s[i]
			count = 1
		}
	}
	sb.WriteString(fmt.Sprintf("%d%s", count, string(c)))
	return sb.String()
}

func solution() (int, int) {
	input := "1321131112"
	for range 40 {
		input = process(input)
	}
	part1 := len(input)
	for range 10 {
		input = process(input)
	}
	part2 := len(input)
	return part1, part2
}

func main() {
	helpers.PrintResult(solution())
}
