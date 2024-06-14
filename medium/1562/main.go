package main

import (
	"strings"
)

func verify(b []byte, l, m int) bool {
	count := 0
	for i := range b {
		if b[i] == '1' {
			count++
		} else if b[i] == '0' {
			if count == m {
				return true
			}
			count = 0
		}
		//log.Printf("i=%d b=%s & count=%d & m=%d", i, b, count, m)
		if l-i < m {
			return count == m
		}
	}
	//log.Printf("b=%s & count=%d & m=%d", b, count, m)
	return count == m
}

func findLatestStep(arr []int, m int) int {
	l := len(arr)
	switch l {
	case 0:
		return -1
	case 1:
		if m == 1 {
			return 1
		}
		return -1
	}
	b := []byte(strings.Repeat("1", l))
	//log.Printf("m=%d & b=%s", l, b)
	for i := l - 1; i > 0; i-- {
		if verify(b, l, m) {
			return i + 1
		}
		b[arr[i]-1] = '0'
	}
	return -1
}

type test struct {
	arr      []int
	m        int
	expected int
}

func main() {
	tests := []test{
		{arr: []int{3, 5, 1, 2, 4}, m: 1, expected: 4},
		{arr: []int{3, 1, 5, 4, 2}, m: 2, expected: -1},
		{arr: []int{1}, m: 1, expected: 1},
		{arr: []int{2, 1}, m: 2, expected: 2},
	}
	for _, t := range tests {
		result := findLatestStep(t.arr, t.m)
		if result != t.expected {
			println("Error: expected", t.expected, "got", result)
		} else {
			println("Success: expected", t.expected, "got", result)
		}
	}
}
