package main

import (
	"log"
	"math"
)

func separate(s []int) int {
	j := 0
	for i, val := range s {
		if val <= 0 {
			s[i] = s[j]
			s[j] = val
			j += 1
		}
	}
	return j
}

func findMissingPositive(s []int) int {
	l := len(s)
	for _, val := range s {
		fval := float64(val)
		if int(math.Abs(fval)-1) < l && s[int(math.Abs(fval)-1)] > 0 {
			s[int(math.Abs(fval)-1)] = -s[int(math.Abs(fval)-1)]
		}
	}

	for i, val := range s {
		if val > 0 {
			return i + 1
		}
	}
	return l + 1
}

func FindMissing(s []int) int {
	shift := separate(s)
	return findMissingPositive(s[shift:])
}

func main() {
	s := []int{3, 4, -1, 1}
	log.Println(FindMissing(s))
	s = []int{1, 2, 0}
	log.Println(FindMissing(s))
}
