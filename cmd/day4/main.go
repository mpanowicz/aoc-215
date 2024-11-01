package main

import (
	"aoc/internal/helpers"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func solution() (int, int) {
	input := "iwrupvqb"

	results := make([]int, 2)
	found5 := false
	found6 := false

	i := 1
	for {
		md := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		s := hex.EncodeToString(md[:])

		if strings.HasPrefix(s, "00000") && !found5 {
			results[0] = i
			found5 = true
		} else if strings.HasPrefix(s, "000000") && !found6 {
			results[1] = i
			found6 = true
		} else if found5 && found6 {
			break
		}

		i++
	}

	return results[0], results[1]
}

func main() {
	helpers.PrintResult(solution())
}
