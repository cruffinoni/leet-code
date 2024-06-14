package main

import "log"

func checkSubarraySum(nums []int, k int) bool {
	l := len(nums)
	if l < 2 {
		return false
	}
	m := make(map[int]int)
	cumulative := 0
	m[0] = -1
	for i := 0; i < l; i++ {
		cumulative += nums[i]
		if k != 0 {
			cumulative = ((cumulative % k) + k) % k
		}
		if idx, ok := m[cumulative]; ok {
			if i-idx >= 2 {
				return true
			}
		} else {
			m[cumulative] = i
		}
	}
	return false
}

type test struct {
	nums     []int
	k        int
	expected bool
}

func main() {
	tests := []test{
		{nums: []int{23, 2, 4, 6, 7}, k: 6, expected: true},
		{nums: []int{23, 2, 6, 4, 7}, k: 6, expected: true},
		{nums: []int{23, 2, 6, 4, 7}, k: 13, expected: false},
		{nums: []int{23, 2, 4, 6, 6}, k: 7, expected: true},
		{nums: []int{5, 0, 0, 0}, k: 3, expected: true},
		{nums: []int{0, 0}, k: 1, expected: true},
		{nums: []int{1, 0}, k: 2, expected: false},
	}
	for _, t := range tests {
		if checkSubarraySum(t.nums, t.k) != t.expected {
			log.Printf("Test failed for %v & %d, expected %t", t.nums, t.k, t.expected)
		} else {
			log.Printf("Test passed for %v", t.nums)
		}
	}
}
