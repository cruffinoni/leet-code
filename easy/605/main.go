package main

import (
	"log"
	"runtime/debug"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	l := len(flowerbed)
	switch l {
	case 0:
		return false
	case 1:
		if n > 1 {
			return false
		}
		return n == 1 && flowerbed[0] == 0
	}
	placed := 0
	for i := 1; i < l; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && (i == 1 || i+1 == l) {
			placed++
			if i == 1 {
				flowerbed[i-1] = 1
			} else if i+1 == l {
				flowerbed[i] = 1
			}
		} else if flowerbed[i-1] == 0 && flowerbed[i] == 0 && i+1 < l && flowerbed[i+1] == 0 {
			placed++
			flowerbed[i] = 1
		}
		if placed >= n {
			return true
		}
	}
	return false
}

func init() {
	debug.SetGCPercent(15)
}

type test struct {
	flowers  []int
	n        int
	expected bool
}

func main() {
	tests := []test{
		{
			flowers:  []int{1, 0, 0, 0, 1},
			n:        1,
			expected: true,
		},
		{
			flowers:  []int{1, 0, 0, 0, 1},
			n:        2,
			expected: false,
		},
		{
			flowers:  []int{0, 0, 1, 0, 1},
			n:        1,
			expected: true,
		},
		{
			flowers:  []int{1, 0, 1, 0, 0},
			n:        1,
			expected: true,
		},
		{
			flowers:  []int{1, 0, 0, 0, 0, 1},
			n:        2,
			expected: false,
		},
		{
			flowers:  []int{0},
			n:        1,
			expected: true,
		},
		{
			flowers:  []int{1},
			n:        0,
			expected: true,
		},
		{
			flowers:  []int{0, 0, 0, 0, 1},
			n:        2,
			expected: true,
		},
	}
	for _, t := range tests {
		cpy := make([]int, len(t.flowers))
		copy(cpy, t.flowers)
		res := canPlaceFlowers(cpy, t.n)
		if res == t.expected {
			log.Printf("test with %v (n=%d) passed!", t.flowers, t.n)
		} else {
			log.Printf("expected %v, got %v w/ %v", t.expected, res, t.flowers)
		}
	}
}
