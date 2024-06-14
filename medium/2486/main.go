package main

import (
	"fmt"
)

func appendCharacters(s string, t string) int {
	l1 := len(s)
	l2 := len(t)
	j := 0
	for i := 0; i < l1 && i < l2; i++ {
		if s[i] == t[j] {
			j++
		}
	}
	return l2 - j
}

type test struct {
	s        string
	t        string
	expected int
}

func main() {
	tests := []test{
		{s: "coaching", t: "coding", expected: 4},
		{s: "abcde", t: "a", expected: 0},
		{s: "z", t: "abcde", expected: 5},
	}
	for _, t := range tests {
		result := appendCharacters(t.s, t.t)
		if result != t.expected {
			fmt.Printf("Error: expected %d, got %d\n", t.expected, result)
		} else {
			fmt.Printf("Success: expected %d, got %d\n", t.expected, result)
		}
	}
}
