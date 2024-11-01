package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
)

func getInput() <-chan string {
	f, _ := os.Open("cmd/day5/input.txt")
	r := bufio.NewReader(f)

	ch := make(chan string)

	go func() {
		for {
			l, _, _ := r.ReadLine()
			s := string(l)
			if s == "" {
				break
			} else {
				ch <- s
			}
		}
		close(ch)
	}()

	return ch
}

func check1(s string) bool {
	vowels := 0
	twice := false
	forbidden := false

	for i, c := range s {
		if forbidden {
			break
		}

		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}

		if i > 0 {
			prev := rune(s[i-1])
			if !twice && prev == c {
				twice = true
			}

			join := string(prev) + string(c)
			switch join {
			case "ab", "cd", "pq", "xy":
				forbidden = true
			}
		}
	}

	return vowels >= 3 && twice && !forbidden
}

func check2(s string) bool {
	dupPair := false
	dupWithOneBetween := false

	dup := make(map[string]bool)
	for i := range s {
		if dupPair && dupWithOneBetween {
			break
		}

		if i > 0 {
			two := s[i-1 : i+1]
			_, ok := dup[two]
			if ok {
				if i > 1 && s[i-2:i] != s[i-1:i+1] {
					dupPair = true
				}
			} else {
				dup[two] = true
			}
		}

		if !dupWithOneBetween && i > 1 && s[i-2] == s[i] {
			dupWithOneBetween = true
		}
	}

	return dupPair && dupWithOneBetween
}

func solution() (int, int) {

	nice1 := 0
	nice2 := 0
	for s := range getInput() {
		if check1(s) {
			nice1++
		}
		if check2(s) {
			nice2++
		}
	}

	return nice1, nice2
}

func main() {
	helpers.PrintResult(solution())
}
