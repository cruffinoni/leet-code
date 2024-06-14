package main

import "log"

func longestWPI(hours []int) int {
	l := len(hours)
	for i := 0; i < l; i++ {
		if hours[i] > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}
}

type test struct {
	hours    []int
	expected int
}

func main() {
	tests := []test{
		{hours: []int{9, 9, 6, 0, 6, 6, 9}, expected: 3},
		{hours: []int{6, 6, 9}, expected: 3},
		{hours: []int{6, 9, 6}, expected: 1},
		{hours: []int{9, 6, 9}, expected: 3},
		{hours: []int{9, 9, 9}, expected: 3},
		{hours: []int{6, 6, 6}, expected: 0},
	}
	for _, t := range tests {
		result := longestWPI(t.hours)
		if result != t.expected {
			log.Printf("Error: expected %v got %v", t.expected, result)
		} else {
			log.Printf("Passed with %v!", t.expected)
		}
	}
}
