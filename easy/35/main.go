package main

import (
	"log"
	"runtime"
)

func searchInsert(nums []int, target int) int {
	runtime.GC()
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
	return lo
}

type test struct {
	nums     []int
	target   int
	expected int
}

func main() {
	tests := []test{
		{nums: []int{1, 3, 5, 6}, target: 5, expected: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, expected: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, expected: 4},
		{nums: []int{1, 3, 5, 6}, target: 0, expected: 0},
	}
	for _, t := range tests {
		res := searchInsert(t.nums, t.target)
		if res != t.expected {
			log.Printf("Test failed for %v, expected %d, got %v", t.nums, t.expected, res)
		} else {
			log.Printf("Test passed for %v", t.nums)
		}
	}
}
