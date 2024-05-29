package main

import (
	"log"
	"math"
)

func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt64, math.MaxInt64

	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			// Found a number greater than 'second'
			return true
		}
	}

	return false
}

type test struct {
	expected bool
	input    []int
}

func main() {
	tests := []test{
		{
			expected: true,
			input:    []int{1, 2, 3, 4, 5},
		},
		{
			expected: false,
			input:    []int{5, 4, 3, 2, 1},
		},
		{
			expected: true,
			input:    []int{2, 1, 5, 0, 4, 6},
		},
		{
			expected: true,
			input:    []int{20, 100, 10, 12, 5, 13},
		},
		{
			expected: true,
			input:    []int{1, 2, 1, 3},
		},
		{
			expected: false,
			input:    []int{1, 2, 2, 1},
		},
		{
			expected: false,
			input:    []int{2, 1, 0, 2, 0, 2, 0, 2, 0, 2, 0, 1},
		},
	}
	for _, t := range tests {
		res := increasingTriplet(t.input)
		if res != t.expected {
			log.Printf("expected %v, got %v for %v", t.expected, res, t.input)
		} else {
			log.Printf("test %v passed!", t.input)
		}
	}
}
