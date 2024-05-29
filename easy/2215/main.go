package main

import (
	"log"
	"slices"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]struct{}, len(nums1))
	m2 := make(map[int]struct{}, len(nums2))
	for _, n := range nums1 {
		m1[n] = struct{}{}
	}
	for _, n := range nums2 {
		m2[n] = struct{}{}
	}
	var result = [][]int{
		make([]int, 0),
		make([]int, 0),
	}
	for n := range m1 {
		if _, ok := m2[n]; !ok {
			result[0] = append(result[0], n)
		}
	}
	for n := range m2 {
		if _, ok := m1[n]; !ok {
			result[1] = append(result[1], n)
		}
	}
	return result
}

type test struct {
	nums1    []int
	nums2    []int
	expected [][]int
}

func main() {
	tests := []test{
		{
			nums1:    []int{1, 2, 3},
			nums2:    []int{2, 4, 6},
			expected: [][]int{{1, 3}, {4, 6}},
		},
	}
	for _, t := range tests {
		result := findDifference(t.nums1, t.nums2)
		if !slices.Equal(result[0], t.expected[0]) || !slices.Equal(result[1], t.expected[1]) {
			log.Printf("Test failed for %v, %v. Got: %v, Expected: %v\n", t.nums1, t.nums2, result, t.expected)
		} else {
			log.Printf("Test passed for %v, %v\n", t.nums1, t.nums2)
		}
	}
}
