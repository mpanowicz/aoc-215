package main

import (
	"aoc/internal/helpers"
	"bufio"
	"bytes"
	"math"
	"os"
	"strconv"
)

type Connection struct {
	From     string
	To       string
	Distance int
}

func getInput() <-chan Connection {
	ch := make(chan Connection)

	go func() {
		f, _ := os.Open("cmd/day9/input.txt")
		r := bufio.NewReader(f)

		for {
			l, _, _ := r.ReadLine()
			if len(l) == 0 {
				break
			}
			parts := bytes.Split(l, []byte(" "))
			distance, _ := strconv.Atoi(string(parts[4]))
			ch <- Connection{string(parts[0]), string(parts[2]), distance}
		}

		close(ch)
	}()

	return ch
}

type Map struct {
	Places map[string]*Place
}

type Place struct {
	Paths map[string]*Path
}

type Path struct {
	Distance    int
	Destination *Place
}

func (m *Map) AddPlace(name string) {
	if _, exists := m.Places[name]; !exists {
		m.Places[name] = &Place{map[string]*Path{}}
	}
}

func (m *Map) AddPath(c *Connection) {
	m.Places[c.From].Paths[c.To] = &Path{c.Distance, m.Places[c.To]}
	m.Places[c.To].Paths[c.From] = &Path{c.Distance, m.Places[c.From]}
}

func swap(places []string, i, j int) {
	temp := places[i]
	places[i] = places[j]
	places[j] = temp
}
func permutation(places []string, idx int) [][]string {
	l := len(places)
	if idx == l {
		tmp := make([]string, l)
		copy(tmp, places)
		return [][]string{tmp}
	}

	result := [][]string{}
	for i := idx; i < l; i++ {
		swap(places, idx, i)
		result = append(result, permutation(places, idx+1)...)
		swap(places, idx, i)
	}

	return result
}
func solution() (int, int) {
	m := &Map{map[string]*Place{}}
	for c := range getInput() {
		m.AddPlace(c.From)
		m.AddPlace(c.To)
		m.AddPath(&c)
	}

	places := []string{}
	for name := range m.Places {
		places = append(places, name)
	}
	routes := permutation(places, 0)

	min := math.MaxInt
	max := math.MinInt
	for _, route := range routes {
		distance := 0
		for i := range len(route) - 1 {
			distance += m.Places[route[i]].Paths[route[i+1]].Distance
		}
		if distance < min {
			min = distance
		}
		if distance > max {
			max = distance
		}
	}

	return min, max
}

func main() {
	helpers.PrintResult(solution())
}
