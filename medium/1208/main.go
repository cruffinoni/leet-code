package main

import (
	"log"
	"math"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func equalSubstring(s string, t string, maxCost int) int {
	start, end := 0, 0
	currentCost := 0
	maxLength := 0

	for end < len(s) {
		currentCost += int(math.Abs(float64(s[end]) - float64(t[end])))
		end++

		for currentCost > maxCost {
			currentCost -= int(math.Abs(float64(s[start]) - float64(t[start])))
			start++
		}

		if end-start > maxLength {
			maxLength = end - start
		}
	}

	return maxLength
}

type test struct {
	s        string
	t        string
	maxCost  int
	expected int
}

func main() {
	tests := []test{
		//{
		//	s:        "abcd",
		//	t:        "bcdf",
		//	maxCost:  3,
		//	expected: 3,
		//},
		//{
		//	s:        "abcd",
		//	t:        "cdef",
		//	maxCost:  3,
		//	expected: 1,
		//},
		//{
		//	s:        "abcd",
		//	t:        "acde",
		//	maxCost:  0,
		//	expected: 1,
		//},
		{
			s:        "krrgw",
			t:        "zjxss",
			maxCost:  19,
			expected: 2,
		},
	}
	for _, t := range tests {
		res := equalSubstring(t.s, t.t, t.maxCost)
		if res != t.expected {
			log.Printf("expected %d, got %d", t.expected, res)
		} else {
			log.Printf("test %s & %s = %d passed!", t.s, t.t, t.maxCost)
		}
	}
}
