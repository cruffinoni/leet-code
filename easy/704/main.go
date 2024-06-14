package main

import "log"

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

type test struct {
	nums     []int
	target   int
	expected int
}

func main() {
	tests := []test{
		//{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, expected: 4},
		//{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, expected: -1},
		//{nums: []int{5}, target: 5, expected: 0},
		//{nums: []int{5}, target: 2, expected: -1},
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 0, expected: 1},
	}
	for _, t := range tests {
		res := search(t.nums, t.target)
		if res != t.expected {
			log.Printf("Test failed for %v, expected %d, got %v", t.nums, t.expected, res)
		} else {
			log.Printf("Test passed for %v", t.nums)
		}
	}
}
