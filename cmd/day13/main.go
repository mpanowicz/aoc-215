package main

import (
	"aoc/internal/helpers"
	"bufio"
	"math"
	"os"
	"strings"
)

type PotentialHappiness struct {
	Name   string
	Value  int
	NextTo string
}

func getInput() <-chan PotentialHappiness {
	ch := make(chan PotentialHappiness)
	go func() {
		f, _ := os.Open("cmd/day13/input.txt")
		r := bufio.NewReader(f)

		for {
			l, _ := r.ReadString('\n')
			if len(l) == 0 {
				break
			}

			split := strings.Split(l[:len(l)-3], " ")
			sign := 0
			switch split[2] {
			case "gain":
				sign = 1
			case "lose":
				sign = -1
			}
			ch <- PotentialHappiness{split[0], sign * helpers.ParseInt(split[3]), split[10]}
		}
		close(ch)
	}()
	return ch
}

type Neighbor struct {
	PotentialHappiness map[string]int
}

type Neighbors map[string]Neighbor

type GetHappiness func(name, next string) int

func calculateHappiness(a []string, gh GetHappiness) int {
	happiness := 0

	l := len(a)
	for i := 0; i < l; i++ {
		var left string
		if i == 0 {
			left = a[l-1]
		} else {
			left = a[i-1]
		}
		happiness += gh(a[i], left)

		var right string
		if i == l-1 {
			right = a[0]
		} else {
			right = a[i+1]
		}
		happiness += gh(a[i], right)
	}

	return happiness
}

func calculate(arrangements [][]string, gh GetHappiness) int {
	max := math.MinInt
	for _, a := range arrangements {
		v := calculateHappiness(a, gh)
		if v > max {
			max = v
		}
	}
	return max
}

const (
	Me = "me"
)

func getConf(me bool, neighbors Neighbors) GetHappiness {
	return func(name, next string) int {
		if me && (name == Me || next == Me) {
			return 0
		}
		return neighbors[name].PotentialHappiness[next]
	}
}

func solution() (int, int) {
	neighbors := Neighbors{}
	names := []string{}

	for ph := range getInput() {
		if n, ok := neighbors[ph.Name]; ok {
			n.PotentialHappiness[ph.NextTo] = ph.Value
		} else {
			neighbors[ph.Name] = Neighbor{map[string]int{ph.NextTo: ph.Value}}
			names = append(names, ph.Name)
		}
	}

	arrangements := helpers.GetPermutations(names)

	p1 := calculate(arrangements, getConf(false, neighbors))
	arrangements = helpers.GetPermutations(append(names, Me))
	p2 := calculate(arrangements, getConf(true, neighbors))

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
