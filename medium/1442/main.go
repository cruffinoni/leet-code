package main

import "log"

func countTriplets(arr []int) int {
	count := 0
	n := len(arr)
	xor := make([]int, n+1)

	for i := 0; i < n; i++ {
		xor[i+1] = xor[i] ^ arr[i]
	}

	freq := make(map[int]int)
	total := make(map[int]int)

	for k := 0; k < n; k++ {
		if val, exists := freq[xor[k+1]]; exists {
			count += val*k - total[xor[k+1]]
		}
		freq[xor[k]]++
		total[xor[k]] += k
	}

	return count
}

type test struct {
	arr      []int
	expected int
}

func main() {
	tests := []test{
		{
			arr:      []int{2, 3, 1, 6, 7},
			expected: 4,
		},
		{
			arr:      []int{1, 1, 1, 1, 1},
			expected: 10,
		},
		{
			arr:      []int{2, 3},
			expected: 0,
		},
	}
	for _, t := range tests {
		result := countTriplets(t.arr)
		if result != t.expected {
			log.Printf("Test failed for %v. Got: %v, Expected: %v\n", t.arr, result, t.expected)
		} else {
			log.Printf("Test passed for %v\n", t.arr)
		}
	}
}
