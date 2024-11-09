package main

import (
	"aoc/internal/helpers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Replacements map[string][]string

type Input struct {
	Replacements Replacements
	Molecule     string
}

func getInput() Input {
	f, _ := os.Open("cmd/day19/input.txt")
	r := bufio.NewReader(f)

	replacementRead := false
	input := Input{
		Replacements: Replacements{},
	}
	for {
		l, _, _ := r.ReadLine()
		if len(l) == 0 {
			if !replacementRead {
				replacementRead = true
				l, _, _ = r.ReadLine()
			} else {
				break
			}
		}

		if !replacementRead {
			parts := strings.Split(string(l), " => ")
			if v, ok := input.Replacements[parts[0]]; ok {
				input.Replacements[parts[0]] = append(v, parts[1])
			} else {
				input.Replacements[parts[0]] = []string{parts[1]}
			}
		} else {
			input.Molecule = string(l)
		}
	}

	return input
}

func replace(s string, position, size int, r Replacements) map[string]struct{} {
	sub := s[position : position+size]

	molecules := map[string]struct{}{}
	if replacements, ok := r[sub]; ok {
		before := ""
		after := ""

		if position > 0 {
			before = s[:position]
		}
		if position+size-1 < len(s)-1 {
			after = s[position+size:]
		}

		for _, r := range replacements {
			molecule := before + r + after
			molecules[molecule] = struct{}{}
		}
	}
	return molecules
}

func getReplaced(input Input, part2 bool) map[string]struct{} {
	replaced := map[string]struct{}{}
	l := len(input.Molecule)

	for i := 0; i < l; i++ {
		if i < l {
			for m := range replace(input.Molecule, i, 1, input.Replacements) {
				replaced[m] = struct{}{}
			}
		}
		if i < l-1 {
			for m := range replace(input.Molecule, i, 2, input.Replacements) {
				replaced[m] = struct{}{}
			}
		}
		fmt.Println(len(replaced))
		if part2 {
			for j := 3; j <= 10; j++ {
				if i < l-j+1 {
					for m := range replace(input.Molecule, i, j, input.Replacements) {
						replaced[m] = struct{}{}
					}
				}
			}
			fmt.Println(len(replaced))
		}
	}
	return replaced
}

func part1(input Input) int {
	replaced := getReplaced(input, false)
	return len(replaced)
}

func part2(s string) int {
	s = strings.ReplaceAll(s, "Rn", "(")
	s = strings.ReplaceAll(s, "Ar", ")")
	s = strings.ReplaceAll(s, "Y", ",")

	fmt.Println(s)

	mol := 0
	par := 0
	com := 0

	for i := len(s) - 1; i >= 0; i-- {
		if 'A' <= s[i] && s[i] <= 'Z' {
			mol++
		} else if s[i] == '(' || s[i] == ')' {
			mol++
			par++
		} else if s[i] == ',' {
			com++
			mol++
		}
	}
	fmt.Println(mol, par, com)

	return mol - par - 2*com - 1
}

func solution() (int, int) {
	input := getInput()

	return part1(input), part2(input.Molecule)
}

func main() {
	helpers.PrintResult(solution())
}
