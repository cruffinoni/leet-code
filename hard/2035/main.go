package main

import (
	"log"
	"math"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func minimumDifference(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	half := totalSum / 2
	log.Printf("%d & %d", n/2+1, half)
	dp := make([][]bool, n/2+1)
	for i := range dp {
		dp[i] = make([]bool, half+1)
		dp[i][0] = true // Somme de 0 est toujours possible
	}

	for i := 1; i <= n; i++ {
		for k := n / 2; k >= 1; k-- {
			for j := half; j >= nums[i-1]; j-- {
				dp[k][j] = dp[k][j] || dp[k-1][j-nums[i-1]]
			}
		}
	}

	for j := half; j >= 0; j-- {
		if dp[n/2][j] {
			sum1 := j
			sum2 := totalSum - sum1
			return int(math.Abs(float64(sum1 - sum2)))
		}
	}

	return -1 // Cela ne devrait jamais arriver
}

type test struct {
	nums     []int
	expected int
}

func main() {
	tests := []test{
		{nums: []int{3, 9, 7, 3}, expected: 2},
		//{nums: []int{-36, 36}, expected: 72},
		{nums: []int{2, -1, 0, 4, -2, -9}, expected: 0},
		//{nums: []int{0, 6, 11, 17, 18, 24}, expected: 6},
	}
	for _, t := range tests {
		result := minimumDifference(t.nums)
		if result != t.expected {
			println("Error: expected", t.expected, "got", result)
		} else {
			println("Success: expected", t.expected, "got", result)
		}
	}
}
