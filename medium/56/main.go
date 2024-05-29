package main

import (
	"log"
	"slices"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	currentInterval := intervals[0]

	for _, interval := range intervals[1:] {
		if currentInterval[1] >= interval[0] {
			currentInterval[1] = max(currentInterval[1], interval[1])
		} else {
			res = append(res, currentInterval)
			currentInterval = interval
		}
	}

	res = append(res, currentInterval)
	return res
}

type test struct {
	input    [][]int
	expected [][]int
}

func main() {
	tests := []test{
		{
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}},
			expected: [][]int{{1, 6}, {8, 10}},
		},
		{
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			input:    [][]int{{1, 4}, {0, 4}},
			expected: [][]int{{0, 4}},
		},
		{
			input:    [][]int{{1, 4}, {0, 0}},
			expected: [][]int{{0, 0}, {1, 4}},
		},
		{
			input:    [][]int{{1, 4}, {0, 2}, {3, 5}},
			expected: [][]int{{0, 5}},
		},
		{
			input:    [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			expected: [][]int{{1, 10}},
		},
	}

	for _, t := range tests {
		out := merge(t.input)
		pass := true
		for i, expected := range t.expected {
			if !slices.Equal(out[i], expected) {
				log.Printf("expected %v, got %v", t.expected, out)
				pass = false
				break
			}
		}
		if pass {
			log.Printf("%v: test passed!", t.expected)
		}
	}
}
