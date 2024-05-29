package main

import (
	"log"
	"sort"
)

func maxOperations(nums []int, k int) int {
	l := len(nums)
	if l == 1 {
		if nums[0] == k {
			return 1
		}
		return 0
	}
	right := l - 1
	left := 0
	sort.Ints(nums)
	count := 0
	for right > left {
		sum := nums[left] + nums[right]
		switch {
		case sum == k:
			count++
			right--
			left++
		case sum > k:
			right--
		case sum < k:
			left++
		}
	}
	return count
}

type test struct {
	nums     []int
	k        int
	expected int
}

func main() {
	tests := []test{
		{
			nums:     []int{1, 2, 3, 4},
			k:        5,
			expected: 2,
		},
		{
			nums:     []int{3, 1, 3, 4, 3},
			k:        6,
			expected: 1,
		},
	}
	for _, t := range tests {
		result := maxOperations(t.nums, t.k)
		if result != t.expected {
			log.Printf("Test failed for %v, %v. Got: %v, Expected: %v\n", t.nums, t.k, result, t.expected)
		} else {
			log.Printf("Test passed for %v, %v\n", t.nums, t.k)
		}
	}
}
