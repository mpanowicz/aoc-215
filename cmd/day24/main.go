package main

import (
	"aoc/internal/helpers"
	"slices"
	"sort"
)

var packages = []int{
	1, 2, 3, 5, 7, 13, 17, 19, 23, 29, 31, 37, 41, 43, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
}

type Group struct {
	Values map[int]struct{}
	Sum    int
}

func generate(subset Group, index int, maxSum int, maxLen int) []Group {
	groups := []Group{}
	groups = append(groups, subset)

	for i := index; i < len(packages); i++ {
		if _, ok := subset.Values[packages[i]]; !ok {
			copy := map[int]struct{}{}
			for k, v := range subset.Values {
				copy[k] = v
			}
			copy[packages[i]] = struct{}{}
			group := Group{copy, (subset.Sum + packages[i])}
			if group.Sum <= maxSum && len(group.Values) <= maxLen {
				groups = append(groups, generate(group, i, maxSum, maxLen)...)
			}
		}
	}

	return groups
}

func quantumEntanglement(g Group) int {
	v := 1
	for k := range g.Values {
		v *= k
	}
	return v
}

func solve(groupSize int) int {
	groups := []Group{}
	n := 2
	for len(groups) == 0 {
		n++
		next := generate(Group{map[int]struct{}{}, 0}, 0, groupSize, n)
		for _, g := range next {
			if g.Sum == groupSize {
				groups = append(groups, g)
			}
		}
	}
	sort.Slice(groups, func(i, j int) bool {
		return quantumEntanglement(groups[i]) < quantumEntanglement(groups[j])
	})

	return quantumEntanglement(groups[0])
}

func solution() (int, int) {
	sum := 0
	slices.Reverse(packages)

	for i := range packages {
		sum += packages[i]
	}

	return solve(sum / 3), solve(sum / 4)
}

func main() {
	helpers.PrintResult(solution())
}
