package main

import "log"

func findMaxAverage(nums []int, k int) float64 {
	l := len(nums)
	if l == 1 {
		return float64(nums[0])
	}
	maxAverage := 0.0
	currentSum := 0
	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}
	maxAverage = float64(currentSum) / float64(k)
	for i := k; i < l; i++ {
		currentSum = currentSum - nums[i-k] + nums[i]
		maxAverage = max(maxAverage, float64(currentSum)/float64(k))
	}
	return maxAverage
}

type test struct {
	nums     []int
	k        int
	expected float64
}

func main() {
	tests := []test{
		{
			nums:     []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75,
		},
		{
			nums:     []int{5},
			k:        1,
			expected: 5,
		},
		{
			nums:     []int{0, 4, 0, 3, 2},
			k:        1,
			expected: 4,
		},
	}
	for _, t := range tests {
		result := findMaxAverage(t.nums, t.k)
		if result != t.expected {
			log.Printf("Test failed for %v, %v. Got: %v, Expected: %v\n", t.nums, t.k, result, t.expected)
		} else {
			log.Printf("Test passed for %v, %v\n", t.nums, t.k)
		}
	}
}
