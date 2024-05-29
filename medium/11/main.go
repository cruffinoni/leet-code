package main

import (
	"log"
	"math"
)

func maxArea(height []int) int {
	l := len(height)
	right := l - 1
	left := 0
	maxSurface := math.MinInt
	for right > left {
		maxSurface = max(maxSurface, min(height[left], height[right])*(right-left))
		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}
	return maxSurface
}

type test struct {
	input    []int
	expected int
}

func main() {
	tests := []test{
		{
			input:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			input:    []int{1, 1},
			expected: 1,
		},
		{
			input:    []int{4, 3, 2, 1, 4},
			expected: 16,
		},
		{
			input:    []int{1, 2, 1},
			expected: 2,
		},
	}
	for _, t := range tests {
		result := maxArea(t.input)
		if result != t.expected {
			log.Printf("Test failed for %v. Got: %v, Expected: %v\n", t.input, result, t.expected)
		} else {
			log.Printf("Test passed for %v\n", t.input)
		}
	}
}
