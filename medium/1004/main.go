package main

import "log"

func longestOnes(nums []int, k int) int {
	l := len(nums)
	if l < k {
		return 0
	}
	left := 0
	right := 0
	currentZeroFlipped := 0
	maxOnes := 0
	for right < l {
		if nums[right] == 0 {
			currentZeroFlipped++
		}
		for currentZeroFlipped > k {
			if nums[left] == 0 {
				currentZeroFlipped--
			}
			left++
		}
		maxOnes = max(right-left+1, maxOnes)
		right++
	}
	return maxOnes
}

type test struct {
	nums     []int
	k        int
	expected int
}

func main() {
	tests := []test{
		//{
		//	nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
		//	k:        2,
		//	expected: 6,
		//},
		//{
		//	nums:     []int{1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
		//	k:        0,
		//	expected: 4,
		//},
		{
			nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:        3,
			expected: 10,
		},
	}
	for _, t := range tests {
		result := longestOnes(t.nums, t.k)
		if result != t.expected {
			log.Printf("Test failed for %v, %v. Got: %v, Expected: %v\n", t.nums, t.k, result, t.expected)
		} else {
			log.Printf("Test passed for %v, %v\n", t.nums, t.k)
		}
	}
}
