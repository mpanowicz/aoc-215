package main

import (
	"aoc/internal/helpers"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

type Gate TokenType

const (
	AND    = "AND"
	OR     = "OR"
	LSHIFT = "LSHIFT"
	RSHIFT = "RSHIFT"
	NOT    = "NOT"

	ARROW = "ARROW"

	WIRE = "WIRE"
	INT  = "INT"
)

var gates = map[string]TokenType{
	"AND":    AND,
	"OR":     OR,
	"LSHIFT": LSHIFT,
	"RSHIFT": RSHIFT,
	"NOT":    NOT,
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) HasNext() bool {
	return l.readPosition <= len(l.input)
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case 0:
		return tok
	case '-':
		l.readChar()
		tok = Token{ARROW, "->"}
	default:
		if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = INT
		} else {
			tok.Literal = l.readText()
			gate, ok := gates[tok.Literal]
			if ok {
				tok.Type = gate
			} else {
				tok.Type = WIRE
			}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readText() string {
	position := l.position
	for l.ch != ' ' && l.ch != 0 {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' {
		l.readChar()
	}
}

type Instruction struct {
	Tokens []Token
}

func getInput() <-chan Instruction {
	ch := make(chan Instruction)
	go func() {
		f, _ := os.Open("cmd/day7/input.txt")
		r := bufio.NewReader(f)

		for {
			tokens := make([]Token, 0)
			l, _, _ := r.ReadLine()
			if len(l) == 0 {
				break
			}

			lex := NewLexer(string(l))
			for lex.HasNext() {
				tokens = append(tokens, lex.NextToken())
			}
			ch <- Instruction{tokens}
		}

		close(ch)
	}()
	return ch
}

func check(calculated *map[string]int, i Instruction) bool {
	ts := i.Tokens
	cal := *calculated
	switch ts[1].Type {
	case WIRE:
		v, ok := cal[ts[1].Literal]
		if ok {
			cal[ts[3].Literal] = ^v
			return true
		} else {
			return false
		}
	case ARROW:
		if ts[0].Type == WIRE {
			v, ok := cal[ts[0].Literal]
			if ok {
				cal[ts[2].Literal] = v
				return true
			} else {
				return false
			}
		} else {
			v, _ := strconv.Atoi(ts[0].Literal)
			cal[ts[2].Literal] = v
			return true
		}
	case LSHIFT:
		v, ok := cal[ts[0].Literal]
		if ok {
			shift, _ := strconv.Atoi(ts[2].Literal)
			cal[ts[4].Literal] = v << shift
			return true
		} else {
			return false
		}
	case RSHIFT:
		v, ok := cal[ts[0].Literal]
		if ok {
			shift, _ := strconv.Atoi(ts[2].Literal)
			cal[ts[4].Literal] = v >> shift
			return true
		} else {
			return false
		}
	case AND, OR:
		var ok bool
		var v1 int
		var v2 int
		if ts[0].Type == INT {
			v1, _ = strconv.Atoi(ts[0].Literal)
		} else {
			v1, ok = cal[ts[0].Literal]
			if !ok {
				return false
			}
		}
		if ts[2].Type == INT {
			v2, _ = strconv.Atoi(ts[2].Literal)
		} else {
			v2, ok = cal[ts[2].Literal]
			if !ok {
				return false
			}
		}
		switch ts[1].Type {
		case AND:
			cal[ts[4].Literal] = v1 & v2
			return true
		case OR:
			cal[ts[4].Literal] = v1 | v2
			return true
		}
	}
	return false
}

func solution() (int, int) {
	instructions := make([]Instruction, 0)

	calculated := make(map[string]int)
	backlog := make([]Instruction, 0)
	for i := range getInput() {
		instructions = append(instructions, i)
		if !check(&calculated, i) {
			backlog = append(backlog, i)
		}
	}

	for len(backlog) > 0 {
		w2 := make([]Instruction, 0)
		for _, i := range backlog {
			if !check(&calculated, i) {
				w2 = append(w2, i)
			}
		}
		backlog = w2
	}
	firstA := calculated["a"]

	calculated = make(map[string]int)
	backlog = instructions
	for len(backlog) > 0 {
		w2 := make([]Instruction, 0)
		for _, i := range backlog {
			if i.Tokens[2].Literal == "b" {
				i.Tokens[0].Literal = fmt.Sprintf("%d", firstA)
			}
			if !check(&calculated, i) {
				w2 = append(w2, i)
			}
		}
		backlog = w2
	}
	secondA := calculated["a"]

	return firstA, secondA
}

func main() {
	helpers.PrintResult(solution())
}
