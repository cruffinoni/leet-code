package main

import (
	"log"
	"slices"
	"sort"
)

func visit(m map[int]int, items [][]int) {
	for _, i := range items {
		m[i[0]] += i[1]
	}
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	maxLen := max(len(items1), len(items2))
	m := make(map[int]int, maxLen)
	visit(m, items1)
	visit(m, items2)
	mapKey := make([]int, 0, maxLen)
	for i := range m {
		mapKey = append(mapKey, i)
	}
	sort.Ints(mapKey)
	res := make([][]int, 0, maxLen)
	for _, i := range mapKey {
		res = append(res, []int{i, m[i]})
	}
	return res
}

type tests struct {
	items1   [][]int
	items2   [][]int
	expected [][]int
}

func main() {
	ts := []tests{
		{
			items1:   [][]int{{1, 1}, {4, 5}, {3, 8}},
			items2:   [][]int{{3, 1}, {1, 5}},
			expected: [][]int{{1, 6}, {3, 9}, {4, 5}},
		},
	}
	for _, t := range ts {
		out := mergeSimilarItems(t.items1, t.items2)
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
