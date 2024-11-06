package main

import (
	"aoc/internal/helpers"
	"aoc/internal/parser"
	"bufio"
	"os"
	"strconv"
)

type Number struct {
	Value    int
	InObject bool
}

func getInput() []parser.JsonToken {
	f, _ := os.Open("cmd/day12/input.txt")
	r := bufio.NewReader(f)
	l, _ := r.ReadString('\n')
	p := parser.New(l)
	tokens := p.ReadObject()
	return tokens
}

func parseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func part1(t *[]parser.JsonToken) int {
	sum := 0
	for _, n := range *t {
		if n.Type == parser.IntValue {
			sum += parseInt(n.Literal)
		}
	}
	return sum
}

func getObjectValue(tokens *[]parser.JsonToken, position int) (sum, nextPosition int) {
	if position >= len(*tokens) {
		return 0, position
	}
	position++
	containsRed := false

loop:
	for {
		t := (*tokens)[position]
		switch t.Type {
		case parser.PropertyName:
			position++
			if (*tokens)[position].Literal == "red" {
				containsRed = true
			}
		case parser.OpenObject:
			s, np := getObjectValue(tokens, position)
			sum += s
			position = np
		case parser.IntValue:
			sum += parseInt(t.Literal)
			position++
		case parser.CloseObject:
			break loop
		default:
			position++
		}
	}

	if containsRed {
		sum = 0
	}
	return sum, position + 1
}

func part2(t *[]parser.JsonToken) int {
	sum, _ := getObjectValue(t, 0)
	return sum
}

func solution() (int, int) {

	tokens := getInput()
	p1 := part1(&tokens)
	p2 := part2(&tokens)

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
