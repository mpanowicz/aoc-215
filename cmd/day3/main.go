package main

import (
	"aoc/internal/helpers"
	"bufio"
	"io"
	"os"
)

type Direction rune

const (
	N Direction = '^'
	S Direction = 'v'
	W Direction = '<'
	E Direction = '>'
)

func getInput() <-chan Direction {
	f, _ := os.Open("cmd/day3/input.txt")
	r := bufio.NewReader(f)

	ch := make(chan Direction)

	go func() {
		for {
			d, _, err := r.ReadRune()
			if err == io.EOF {
				break
			} else {
				switch d {
				case '^':
					ch <- N
				case 'v':
					ch <- S
				case '<':
					ch <- W
				case '>':
					ch <- E
				}
			}
		}
		close(ch)
	}()

	return ch
}

type Position struct {
	x int
	y int
}

func NewPosition(x, y int) *Position {
	return &Position{x, y}
}

func (p *Position) Move(d Direction) {
	switch d {
	case N:
		p.x += 1
	case S:
		p.x -= 1
	case W:
		p.y -= 1
	case E:
		p.y += 1
	}
}

type Visited map[Position]int

func (v Visited) UpdateVisited(p *Position) {
	val, ok := v[*p]
	if ok {
		v[*p] = val + 1
	} else {
		v[*p] = 1
	}
}

type Santa struct {
	position *Position
	visited  *Visited
}

func NewSanta(v *Visited) Santa {
	s := Santa{NewPosition(0, 0), v}
	s.visited.UpdateVisited(s.position)
	return s
}

func (s Santa) Move(d Direction) {
	s.position.Move(d)
	s.visited.UpdateVisited(s.position)
}

func solution(numOfSanta int) int {
	visited := make(Visited)

	santas := make([]Santa, 0)
	for i := 0; i < numOfSanta; i++ {
		santas = append(santas, NewSanta(&visited))
	}

	sIndex := 0
	for d := range getInput() {
		santa := santas[sIndex%numOfSanta]
		santa.Move(d)
		sIndex++
	}
	numberOfVisited := len(visited)

	return numberOfVisited
}

func main() {
	helpers.PrintResult(solution(1), solution(2))
}
