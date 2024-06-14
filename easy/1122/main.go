package main

import (
	"log"
	"reflect"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	// Step 1: Create a frequency map for arr1
	freq := make(map[int]int)
	for _, num := range arr1 {
		freq[num]++
	}

	// Step 2: Initialize the result array
	result := []int{}

	// Step 3: Add elements to result based on the order in arr2
	for _, num := range arr2 {
		if count, exists := freq[num]; exists {
			for i := 0; i < count; i++ {
				result = append(result, num)
			}
			delete(freq, num) // Remove the element from the frequency map once added
		}
	}

	// Step 4: Collect remaining elements
	remaining := []int{}
	for num, count := range freq {
		for i := 0; i < count; i++ {
			remaining = append(remaining, num)
		}
	}

	// Step 5: Sort remaining elements and append to result
	sort.Ints(remaining)
	result = append(result, remaining...)

	return result
}

type test struct {
	arr1     []int
	arr2     []int
	expected []int
}

func main() {
	tests := []test{
		{arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, arr2: []int{2, 1, 4, 3, 9, 6}, expected: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}},
	}
	for _, t := range tests {
		res := relativeSortArray(t.arr1, t.arr2)
		if !reflect.DeepEqual(res, t.expected) {
			log.Printf("Test failed for %v, expected %v, got %v", t.arr1, t.expected, res)
		} else {
			log.Printf("Test passed for %v", t.arr1)
		}
	}
}
