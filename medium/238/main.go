package main

import (
	"log"
	"slices"
)

func productExceptSelf(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	right := make([]int, l)
	left := make([]int, l)
	k := 1
	left[0] = 1
	for k < l {
		left[k] = left[k-1] * nums[k-1]
		k++
	}
	k = l - 2
	right[l-1] = 1
	for k >= 0 {
		right[k] = right[k+1] * nums[k+1]
		k--
	}

	log.Printf("Left: %v & Right: %v", left, right)
	for i := 0; i < l; i++ {
		res[i] = left[i] * right[i]
	}
	return res
}

type test struct {
	input    []int
	expected []int
}

func main() {
	tests := []test{
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		//{
		//	input:    []int{-1, 1, 0, -3, 3},
		//	expected: []int{0, 0, 9, 0, 0},
		//},
	}
	for _, t := range tests {
		out := productExceptSelf(t.input)
		if !slices.Equal(out, t.expected) {
			log.Printf("expected %v, got %v", t.expected, out)
		} else {
			log.Printf("%v: test passed!", t.expected)
		}
	}
}
