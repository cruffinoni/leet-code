package main

import (
	"log"
	"slices"
)

func singleNumber(nums []int) []int {
	sum := 0
	for i := range nums {
		sum ^= nums[i]
	}
	separator := 1
	for i := 0; separator&sum == 0; i++ {
		separator = 1 << i
	}
	uniqueA, uniqueB := 0, 0
	for i := range nums {
		if nums[i]&separator > 0 {
			uniqueA ^= nums[i]
		} else {
			uniqueB ^= nums[i]
		}
	}

	return []int{uniqueA, uniqueB}
}

type test struct {
	nums     []int
	expected []int
}

func main() {
	tests := []test{
		{
			nums:     []int{1, 2, 1, 3, 2, 5},
			expected: []int{3, 5},
		},
		{
			nums:     []int{2, 1, 2, 3, 4, 1},
			expected: []int{3, 4},
		},
	}
	for _, t := range tests {
		result := singleNumber(t.nums)
		if !slices.Equal(result, t.expected) {
			log.Printf("Test failed for %v. Got: %v, Expected: %v\n", t.nums, result, t.expected)
		} else {
			log.Printf("Test passed for %v\n", t.nums)
		}
	}
}
