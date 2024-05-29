package main

import (
	"log"
	"sort"
)

func specialArray(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i <= n; i++ {
		// On cherche le nombre de valeurs >= i
		pos := sort.Search(n, func(j int) bool {
			return nums[j] >= i
		})
		if n-pos == i {
			return i
		}
	}
	return -1
}

func main() {
	tests := map[int][]int{
		2:  {3, 5},
		-1: {0, 0},
		3:  {0, 4, 3, 0, 4},
	}

	for expected, input := range tests {
		res := specialArray(input)
		if res != expected {
			log.Printf("expected %d, got %d with input %v", expected, res, input)
		} else {
			log.Printf("test with input %v passed!", input)
		}
	}
}
