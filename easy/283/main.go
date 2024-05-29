package main

import (
	"log"
	"slices"
)

func moveZeroes(nums []int) {
	l := len(nums)
	if l == 1 {
		return
	}
	nonZeroIndex := 0
	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex] = nums[i]
			nonZeroIndex++
		}
	}
	for nonZeroIndex < l {
		nums[nonZeroIndex] = 0
		nonZeroIndex++
	}
}

type test struct {
	in       []int
	expected []int
}

func main() {
	tests := []test{
		{
			in:       []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			in:       []int{0, 1, 0, 3},
			expected: []int{1, 3, 0, 0},
		},
		{
			in:       []int{0, 0, 1, 3},
			expected: []int{1, 3, 0, 0},
		},
		{
			in:       []int{0, 1},
			expected: []int{1, 0},
		},
		{
			in:       []int{0},
			expected: []int{0},
		},
	}
	for _, t := range tests {
		cpyIn := make([]int, len(t.in))
		copy(cpyIn, t.in)
		moveZeroes(cpyIn)
		if !slices.Equal(t.expected, cpyIn) {
			log.Printf("expected %v, got %v for %v", t.expected, cpyIn, t.in)
		} else {
			log.Printf("test %v passed!", t.in)
		}
	}
}
