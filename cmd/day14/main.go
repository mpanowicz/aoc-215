package main

import (
	"aoc/internal/helpers"
	"bufio"
	"math"
	"os"
	"strings"
)

type Reindeer struct {
	Name      string
	Speed     int
	Endurance int
	Rest      int
}

func (r *Reindeer) SegmentTime() int {
	return r.Endurance + r.Rest
}

func (r *Reindeer) FlySegmentDistance() int {
	return r.Speed * r.Endurance
}

func (r *Reindeer) DistanceAfterTime(time int) int {
	st := r.SegmentTime()
	segments := time / st
	lastPart := time % st
	distance := segments * r.Speed * r.Endurance
	if lastPart < r.Endurance {
		distance += lastPart * r.Speed
	} else {
		distance += r.FlySegmentDistance()
	}
	return distance
}

func getInput() <-chan Reindeer {
	ch := make(chan Reindeer)
	go func() {
		f, _ := os.Open("cmd/day14/input.txt")
		r := bufio.NewReader(f)

		for {
			l, _ := r.ReadString('\n')
			if len(l) == 0 {
				break
			}

			line := string(l[:len(l)-1])
			parts := strings.Split(line, " ")
			ch <- Reindeer{
				parts[0],
				helpers.ParseInt(parts[3]),
				helpers.ParseInt(parts[6]),
				helpers.ParseInt(parts[13]),
			}
		}

		close(ch)
	}()
	return ch
}

func solution() (int, int) {
	checkTime := 2503
	reindeers := []Reindeer{}

	part1 := math.MinInt
	for r := range getInput() {
		reindeers = append(reindeers, r)
		distance := r.DistanceAfterTime(checkTime)
		if distance > part1 {
			part1 = distance
		}
	}

	pointsBoard := map[string]int{}
	for i := 1; i <= checkTime; i++ {
		intervalMax := math.MinInt
		var name string
		for _, r := range reindeers {
			distance := r.DistanceAfterTime(i)
			if distance > intervalMax {
				intervalMax = distance
				name = r.Name
			}
		}

		pointsBoard[name] += 1
	}
	part2 := math.MinInt
	for _, points := range pointsBoard {
		if points > part2 {
			part2 = points
		}
	}

	return part1, part2
}

func main() {
	helpers.PrintResult(solution())
}
