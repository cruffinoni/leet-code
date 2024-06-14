package main

import (
	"log"
	"slices"
)

func sortColors(nums []int) {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	res := make([]int, 0, len(nums))
	for i := 0; i < 3; i++ {
		for j := 0; j < m[i]; j++ {
			res = append(res, i)
		}
	}
	copy(nums, res)
}

type test struct {
	nums     []int
	expected []int
}

func main() {
	tests := []test{
		{nums: []int{2, 0, 2, 1, 1, 0}, expected: []int{0, 0, 1, 1, 2, 2}},
		{nums: []int{2, 0, 1}, expected: []int{0, 1, 2}},
		{nums: []int{0}, expected: []int{0}},
		{nums: []int{1}, expected: []int{1}},
	}
	for _, t := range tests {
		cpy := make([]int, len(t.nums))
		copy(cpy, t.nums)
		sortColors(t.nums)
		if !slices.Equal(t.nums, t.expected) {
			log.Printf("Test failed for %v, expected %v, got %v", cpy, t.expected, t.nums)
		} else {
			log.Printf("Test passed for %v", t.nums)
		}
	}
}
