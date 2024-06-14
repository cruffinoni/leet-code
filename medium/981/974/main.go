package main

import "log"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func subarraysDivByK(nums []int, k int) int {
	remainderFreq := make(map[int]int)
	prefixSum := 0
	count := 0
	remainderFreq[0] = 1
	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		idx := abs(prefixSum % k)
		remainderFreq[idx]++
		count += remainderFreq[idx]
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
		{nums: []int{4, 5, 0, -2, -3, 1}, k: 5, expected: 7},
		//{nums: []int{-1, 2, 9}, k: 2, expected: 2},
	}
	for _, t := range tests {
		res := subarraysDivByK(t.nums, t.k)
		if res != t.expected {
			log.Printf("Test failed for %v, expected %d, got %v", t.nums, t.expected, res)
		} else {
			log.Printf("Test passed for %v", t.nums)
		}
	}
}
