package main

import (
	"aoc/internal/helpers"

	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
)

type Box struct {
	l                int
	w                int
	h                int
	lw               int
	wh               int
	hl               int
	smallestSurface  int
	smallestSidesSum int
}

func NewBox(l, w, h int) Box {
	b := Box{
		l:  l,
		w:  w,
		h:  h,
		lw: l * w,
		wh: w * h,
		hl: h * l,
	}

	b.smallestSurface = b.SmallestSurface()
	b.smallestSidesSum = b.SmallestSidesSum()

	return b
}

func (b Box) SmallestSurface() int {
	if b.lw < b.wh {
		if b.lw < b.hl {
			return b.lw
		} else {
			return b.hl
		}
	} else {
		if b.wh < b.hl {
			return b.wh
		} else {
			return b.hl
		}
	}
}

func (b Box) SmallestSidesSum() int {
	if b.l > b.w {
		if b.l > b.h {
			return b.w + b.h
		} else {
			return b.l + b.w
		}
	} else {
		if b.w > b.h {
			return b.l + b.h
		} else {
			return b.l + b.w
		}
	}
}

func (b Box) Wrapping() int {
	return 2*(b.lw+b.wh+b.hl) + b.smallestSurface
}

func (b Box) Ribbon() int {
	return 2*b.smallestSidesSum + b.l*b.w*b.h
}

func toInt(b []byte) int {
	i, _ := strconv.Atoi(string(b))
	return i
}

func getInput() <-chan Box {
	f, _ := os.Open("cmd/day2/input.txt")
	r := bufio.NewReader(f)

	ch := make(chan Box)

	go func() {
		for {
			line, _, err := r.ReadLine()
			if err == io.EOF || len(line) == 0 {
				break
			} else {
				parts := bytes.Split(line, []byte("x"))

				l, w, h := toInt(parts[0]), toInt(parts[1]), toInt(parts[2])

				ch <- NewBox(l, w, h)
			}
		}
		close(ch)
	}()

	return ch
}

func solution() (int, int) {

	wrapping := 0
	ribbon := 0
	for b := range getInput() {
		wrapping += b.Wrapping()
		ribbon += b.Ribbon()
	}

	return wrapping, ribbon
}

func main() {
	helpers.PrintResult(solution())
}
