package helpers

import (
	"fmt"
	"strconv"
)

func ParseInt(s string) int {
	v, e := strconv.Atoi(s)
	if e != nil {
		fmt.Println(e)
	}
	return v
}

func SumInt(l []int) int {
	s := 0
	for i := range l {
		s += l[i]
	}
	return s
}
