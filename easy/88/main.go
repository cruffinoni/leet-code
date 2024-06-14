package main

import (
	"log"
	"slices"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	} else if n == 0 {
		return
	}
	log.Printf("nums1: %v, nums2: %v", nums1, nums2)
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}

type test struct {
	nums1    []int
	m        int
	nums2    []int
	n        int
	expected []int
}

func main() {
	tests := []test{
		{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, expected: []int{1, 2, 2, 3, 5, 6}},
	}
	for _, t := range tests {
		cpyNums1 := make([]int, len(t.nums1))
		copy(cpyNums1, t.nums1)
		merge(t.nums1, t.m, t.nums2, t.n)
		if !slices.Equal(t.nums1, t.expected) {
			log.Printf("Test failed for %v & %v, expected %v, got %v", cpyNums1, t.nums2, t.expected, t.nums1)
		} else {
			log.Printf("Test passed for %v", t.nums1)
		}
	}
}
