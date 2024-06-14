package main

import (
	"log"
	"sort"
)

func minIncrementForUnique(nums []int) int {
	count := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			count += (nums[i] + 1) - nums[i+1]
			nums[i+1] = nums[i] + 1
		}
	}
	return count
}

type test struct {
	nums     []int
	expected int
}

func main() {
	tests := []test{
		//{nums: []int{1, 2, 2}, expected: 1},
		{nums: []int{3, 2, 1, 2, 1, 7}, expected: 6},
	}
	for _, t := range tests {
		res := minIncrementForUnique(t.nums)
		if res != t.expected {
			log.Printf("Test failed for %v, expected %d, got %v", t.nums, t.expected, res)
		} else {
			log.Printf("Test passed for %v", t.nums)
		}
	}
}
