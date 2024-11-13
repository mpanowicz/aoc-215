package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
	"strings"
)

const (
	IncrementA = "inc a"
	IncrementB = "inc b"
	Triple     = "tpl"
	Half       = "hlf"
	Jump       = "jmp"
	JumpEven   = "jie"
	JumpOne    = "jio"
)

type Action struct {
	Type   string
	Offset int
}

func Apply(r1, r2 int, a Action) (int, int, int) {
	switch a.Type {
	case IncrementA:
		return r1 + 1, r2, 1
	case IncrementB:
		return r1, r2 + 1, 1
	case Triple:
		return 3 * r1, r2, 1
	case Half:
		return r1 / 2, r2, 1
	case Jump:
		return r1, r2, a.Offset
	case JumpEven:
		if r1%2 == 0 {
			return r1, r2, a.Offset
		} else {
			return r1, r2, 1
		}
	case JumpOne:
		if r1 == 1 {
			return r1, r2, a.Offset
		} else {
			return r1, r2, 1
		}
	}
	return r1, r2, 1
}

func getInput() []Action {
	actions := []Action{}

	f, _ := os.Open("cmd/day23/input.txt")
	r := bufio.NewReader(f)

	for {
		l, _, _ := r.ReadLine()
		if len(l) == 0 {
			break
		}
		line := string(l)

		if strings.HasPrefix(line, IncrementA) {
			actions = append(actions, Action{IncrementA, 1})
		} else if strings.HasPrefix(line, IncrementB) {
			actions = append(actions, Action{IncrementB, 1})
		} else if strings.HasPrefix(line, Triple) {
			actions = append(actions, Action{Triple, 1})
		} else if strings.HasPrefix(line, Half) {
			actions = append(actions, Action{Half, 1})
		} else if strings.HasPrefix(line, Jump) {
			offset := helpers.ParseInt(line[len(Jump+" "):])
			actions = append(actions, Action{Jump, offset})
		} else if strings.HasPrefix(line, JumpEven) {
			offset := helpers.ParseInt(line[len(JumpEven+" a, "):])
			actions = append(actions, Action{JumpEven, offset})
		} else if strings.HasPrefix(line, JumpOne) {
			offset := helpers.ParseInt(line[len(JumpOne+" a, "):])
			actions = append(actions, Action{JumpOne, offset})
		} else {
			panic("aaa")
		}
	}

	return actions
}

func compute(a int, actions []Action) int {
	offset := 0
	b := 0
	for 0 <= offset && offset < len(actions) {
		aChange, bChange, offsetChange := Apply(a, b, actions[offset])
		a = aChange
		b = bChange
		offset += offsetChange
	}

	return b
}

func solution() (int, int) {
	actions := getInput()

	return compute(0, actions), compute(1, actions)
}

func main() {
	helpers.PrintResult(solution())
}
