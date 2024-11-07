package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
	"strings"
)

const (
	Children    = "children"
	Cats        = "cats"
	Samoyeds    = "samoyeds"
	Pomeranians = "pomeranians"
	Akitas      = "akitas"
	Vizslas     = "vizslas"
	Goldfish    = "goldfish"
	Trees       = "trees"
	Cars        = "cars"
	Perfumes    = "perfumes"
)

type Gifts map[string]int

type Aunt struct {
	Number int
	Gifts  Gifts
}

func getInput() <-chan Aunt {
	ch := make(chan Aunt)
	go func() {
		f, _ := os.Open("cmd/day16/input.txt")
		r := bufio.NewReader(f)

		for {
			l, _, _ := r.ReadLine()
			if len(l) == 0 {
				break
			}

			line := string(l)
			nameSplit := strings.Split(line, ": ")
			name := strings.Split(nameSplit[0], " ")[1]
			nameId := helpers.ParseInt(name)
			elements := strings.Split(line[len(nameSplit[0])+2:], ", ")
			gifts := Gifts{}
			for _, e := range elements {
				parts := strings.Split(e, ": ")
				gifts[parts[0]] = helpers.ParseInt(parts[1])
			}
			ch <- Aunt{nameId, gifts}
		}

		close(ch)
	}()
	return ch
}

type Check func(expected, current int) bool

func solution() (int, int) {
	part2conf := map[string]Check{
		Cats:        func(e, c int) bool { return e < c },
		Trees:       func(e, c int) bool { return e < c },
		Pomeranians: func(e, c int) bool { return e > c },
		Goldfish:    func(e, c int) bool { return e > c },
	}
	readings := Gifts{
		Children:    3,
		Cats:        7,
		Samoyeds:    2,
		Pomeranians: 3,
		Akitas:      0,
		Vizslas:     0,
		Goldfish:    5,
		Trees:       3,
		Cars:        2,
		Perfumes:    1,
	}
	part1 := 0
	part2 := 0
	for a := range getInput() {
		found1 := true
		found2 := true
		for n, v := range a.Gifts {
			if readings[n] != v {
				found1 = found1 && false
			}
			if check, ok := part2conf[n]; ok {
				if !check(readings[n], v) {
					found2 = found2 && false
				}
			} else if readings[n] != v {
				found2 = found2 && false
			}
		}
		if found1 {
			part1 = a.Number
		}
		if found2 {
			part2 = a.Number
		}
	}

	return part1, part2
}

func main() {
	helpers.PrintResult(solution())
}
