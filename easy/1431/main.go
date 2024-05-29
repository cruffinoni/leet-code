package main

import (
	"log"
	"slices"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := slices.Max(candies)
	l := len(candies)
	res := make([]bool, l)
	for i := 0; i < l; i++ {
		res[i] = candies[i]+extraCandies >= maxCandies
	}
	return res
}

type test struct {
	candies      []int
	extraCandies int
	expected     []bool
}

func main() {
	tests := []test{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			expected:     []bool{true, true, true, false, true},
		},
		{
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			expected:     []bool{true, false, false, false, false},
		},
	}
	for _, t := range tests {
		res := kidsWithCandies(t.candies, t.extraCandies)
		if slices.Equal(res, t.expected) {
			log.Printf("test with %v (extra=%d) passed!", t.candies, t.extraCandies)
		} else {
			log.Printf("expected %v, got %v", t.expected, res)
		}
	}
}
