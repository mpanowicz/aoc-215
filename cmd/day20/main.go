package main

import (
	"aoc/internal/helpers"
)

func getDivisors(n int) []int {
	divisors := []int{}

	for i := 1; i*i <= n; i++ {
		if i*i == n {
			divisors = append(divisors, i)
		} else {
			if n%i == 0 {
				divisors = append(divisors, []int{i, n / i}...)
			}
		}
	}

	return divisors
}

type Result struct {
	Part  int
	House int
}

func part1(input int, ch chan Result) {
	house := 1
	for {
		sum := 0
		for _, d := range getDivisors(house) {
			sum += d
		}
		if sum*10 >= input {
			break
		}
		house++
	}
	ch <- Result{1, house}
}

func part2(input int, ch chan Result) {
	house := 1
	for {
		sum := 0
		for _, d := range getDivisors(house) {
			if house/d <= 50 {
				sum += d
			}
		}
		if sum*11 >= input {
			break
		}
		house++
	}
	ch <- Result{2, house}
}

func solution() (int, int) {
	input := 33100000
	ch := make(chan Result)
	go func() {
		part1(input, ch)
	}()
	go func() {
		part2(input, ch)
	}()

	p1 := 0
	p2 := 0

	for r := range ch {
		if r.Part == 1 {
			p1 = r.House
		} else {
			p2 = r.House
		}
		if p1 != 0 && p2 != 0 {
			break
		}
	}
	close(ch)

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
