package main

import (
	"log"
	"slices"
	"sort"
)

const maxResSize = 2

func findWinners(matches [][]int) [][]int {
	winners := make(map[int]int, len(matches))
	loosers := make(map[int]int, len(matches))
	for _, i := range matches {
		winners[i[0]]++
		loosers[i[1]]++
	}
	res := make([][]int, maxResSize)
	for w := range winners {
		if _, ok := loosers[w]; !ok {
			res[0] = append(res[0], w)
		}
	}
	for id, l := range loosers {
		if l == 1 {
			res[1] = append(res[1], id)
		}
	}
	for i := 0; i < maxResSize; i++ {
		sort.Ints(res[i])
	}
	return res
}

type test struct {
	input    [][]int
	expected [][]int
}

func main() {
	tests := []test{
		{
			input:    [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}},
			expected: [][]int{{1, 2, 10}, {4, 5, 7, 8}},
		},
	}
	for _, t := range tests {
		out := findWinners(t.input)
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
