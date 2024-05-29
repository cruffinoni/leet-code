package main

import "log"

func largestAltitude(gain []int) int {
	currentMax := 0
	currentAltitude := 0
	for _, i := range gain {
		currentAltitude += i
		currentMax = max(currentMax, currentAltitude)
	}
	return currentMax
}

type test struct {
	gain     []int
	expected int
}

func main() {
	tests := []test{
		{
			gain:     []int{-5, 1, 5, 0, -7},
			expected: 1,
		},
		{
			gain:     []int{-4, -3, -2, -1, 4, 3, 2},
			expected: 0,
		},
	}
	for _, t := range tests {
		result := largestAltitude(t.gain)
		if result != t.expected {
			log.Printf("Test failed for %v. Got: %v, Expected: %v\n", t.gain, result, t.expected)
		} else {
			log.Printf("Test passed for %v\n", t.gain)
		}
	}
}
