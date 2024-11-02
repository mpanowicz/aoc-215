package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type InstructionType string

const (
	On     InstructionType = "on"
	Off    InstructionType = "off"
	Toggle InstructionType = "toggle"
)

type Point struct {
	Row int
	Col int
}

type Instruction struct {
	Type  InstructionType
	Start Point
	End   Point
}

type change func(int) int
type Configuration struct {
	onOn     change
	onOff    change
	onToggle change
}

func parseLine(l string) Instruction {
	iType, left := getType(l)
	s, e := getRange(left)
	return Instruction{iType, s, e}
}

func getType(l string) (InstructionType, string) {
	if left, found := strings.CutPrefix(l, "turn on "); found {
		return On, left
	} else if left, found := strings.CutPrefix(l, "turn off "); found {
		return Off, left
	} else {
		left, _ := strings.CutPrefix(l, "toggle ")
		return Toggle, left
	}
}

func getRange(l string) (Point, Point) {
	parts := strings.Split(l, " through ")
	return getPoint(parts[0]), getPoint(parts[1])
}

func getPoint(l string) Point {
	parts := strings.Split(l, ",")
	r, _ := strconv.Atoi(parts[0])
	c, _ := strconv.Atoi(parts[1])
	return Point{r, c}
}

func getInput() <-chan Instruction {
	f, _ := os.Open("cmd/day6/input.txt")
	r := bufio.NewReader(f)

	ch := make(chan Instruction)
	go func() {
		for {
			l, _, _ := r.ReadLine()
			line := string(l)
			if line == "" {
				break
			}
			ch <- parseLine(line)
		}

		close(ch)
	}()
	return ch
}

type Grid [1000][1000]int

type Lights struct {
	Grid          *Grid
	Configuration Configuration
}

func (l Lights) Configure(i Instruction) {
	switch i.Type {
	case On:
		l.On(i)
	case Off:
		l.Off(i)
	case Toggle:
		l.Toggle(i)
	}
}

func (l Lights) On(i Instruction) {
	l.Grid.changeLight(i, l.Configuration.onOn)
}
func (l Lights) Off(i Instruction) {
	l.Grid.changeLight(i, l.Configuration.onOff)
}
func (l Lights) Toggle(i Instruction) {
	l.Grid.changeLight(i, l.Configuration.onToggle)
}

func (g *Grid) changeLight(i Instruction, fn change) {
	for r := i.Start.Row; r <= i.End.Row; r++ {
		for c := i.Start.Col; c <= i.End.Col; c++ {
			g[r][c] = fn(g[r][c])
		}
	}
}

func solution(c Configuration) int {
	g := Grid([1000][1000]int{})
	lights := Lights{&g, c}
	for i := range getInput() {
		lights.Configure(i)
	}

	on := 0
	for _, r := range g {
		for _, v := range r {
			on += v
		}
	}

	return on
}

func main() {
	c1 := Configuration{
		func(v int) int { return 1 },
		func(v int) int { return 0 },
		func(v int) int { return 1 - v },
	}
	c2 := Configuration{
		func(v int) int { return v + 1 },
		func(v int) int { return max(0, v-1) },
		func(v int) int { return v + 2 },
	}

	helpers.PrintResult(solution(c1), solution(c2))
}
