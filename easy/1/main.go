package main

import (
	"log"
	"slices"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if idx, ok := m[target-n]; ok {
			return []int{idx, i}
		}
		m[n] = i
	}
	return nil
}

type test struct {
	nums     []int
	target   int
	expected []int
}

func main() {
	tests := []test{
		{nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
	}
	for _, t := range tests {
		if !slices.Equal(twoSum(t.nums, t.target), t.expected) {
			log.Printf("Test failed for %v & %d, expected %v", t.nums, t.target, t.expected)
		} else {
			log.Printf("Test passed for %v", t.nums)
		}
	}
}
