package main

import (
	"log"
	"reflect"
)

func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}

	return ans
}

/*
0 (0) = 0
1 (1) = 0

10 (2) = 1
11 (3) = 1

100 (4) = 2
101 (5) = 2
110 (6) = 3
111 (7) = 3

1000 (8) = 1
1001 (9) = 2
1010 (10) = 2
1011 (11) = 3
1100 (12) = 2
1101 (13) = 3
1110 (14) = 3
1111 (15) = 4
*/

type test struct {
	n        int
	expected []int
}

func main() {
	tests := []test{
		//{
		//	n:        2,
		//	expected: []int{0, 1, 1},
		//},
		{
			n:        5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, t := range tests {
		result := countBits(t.n)
		if !reflect.DeepEqual(result, t.expected) {
			log.Printf("Test failed for %v. Got: %v, Expected: %v\n", t.n, result, t.expected)
		} else {
			log.Printf("Test passed for %v\n", t.n)
		}
	}
}
