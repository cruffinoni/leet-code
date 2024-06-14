package main

import "log"

func getConsecutive(m map[int]struct{}, n int) int {
	if _, ok := m[n]; !ok {
		return 0
	}
	for i := n + 1; ; i++ {
		if _, ok := m[i]; !ok {
			return i - n
		}
	}
}

func longestConsecutive(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	set := make(map[int]struct{})
	for i := 0; i < l; i++ {
		set[nums[i]] = struct{}{}
	}
	m := make(map[int]struct{})
	maxSequence := 0
	for i := 0; i < l; i++ {
		m[nums[i]] = struct{}{}
		if _, ok := set[nums[i]-1]; !ok {
			maxSequence = max(getConsecutive(set, nums[i]), maxSequence)
		}
	}
	return maxSequence
}

type test struct {
	nums     []int
	expected int
}

func main() {
	tests := []test{
		{nums: []int{100, 4, 200, 1, 3, 2}, expected: 4},
		{nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, expected: 9},
	}
	for _, t := range tests {
		res := longestConsecutive(t.nums)
		if res != t.expected {
			log.Printf("Test failed for %v, expected %d, go %v", t.nums, t.expected, res)
		} else {
			log.Printf("Test passed for %v", t.nums)
		}
	}
}
