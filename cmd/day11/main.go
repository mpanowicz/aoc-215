package main

import (
	"aoc/internal/helpers"
)

func hasThree(s []byte) bool {
	for i := 1; i < len(s)-1; i++ {
		v1 := s[i-1]
		v2 := s[i]
		v3 := s[i+1]
		if v1+1 == v2 && v2+1 == v3 {
			return true
		}
	}
	return false
}
func correctPassword(s []byte) bool {
	pairs := map[byte]bool{}
	straight := hasThree(s)

	for i, c := range s {
		if c == 'i' || c == 'o' || c == 'l' {
			return false
		}
		if len(pairs) < 2 {
			if i < len(s)-1 {
				if !pairs[c] && c == s[i+1] {
					pairs[c] = true
				}
				i++
			}
		}
	}
	return len(pairs) >= 2 && straight
}

func nextPassword(s []byte) {
	update := true
	pos := len(s) - 1
	for update {
		c := s[pos] + 1
		if c > 'z' {
			s[pos] = 'a' + ((c - 'a') % 26)
			pos--
		} else {
			s[pos] = c
			update = false
		}
	}
}

func getCorrect(s string) string {
	bytes := []byte(s)
	for !correctPassword(bytes) {
		nextPassword(bytes)
	}
	return string(bytes)
}

func solution() (string, string) {
	input := "cqjxjnds"
	part1 := getCorrect(input)

	next := []byte(part1)
	nextPassword(next)
	input = string(next)
	part2 := getCorrect(input)

	return part1, part2
}

func main() {
	helpers.PrintResult(solution())
}
