package main

import "log"

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	idxS := 0
	l := len(s)
	for i := 0; i < len(t); i++ {
		if s[idxS] == t[i] {
			idxS++
			if idxS == l {
				return true
			}
		}
	}
	return false
}

type test struct {
	s        string
	t        string
	expected bool
}

func main() {
	tests := []test{
		{
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			s:        "axc",
			t:        "ahbgdc",
			expected: false,
		},
	}
	for _, t := range tests {
		result := isSubsequence(t.s, t.t)
		if result != t.expected {
			log.Printf("Test failed for %v, %v. Got: %v, Expected: %v\n", t.s, t.t, result, t.expected)
		} else {
			log.Printf("Test passed for %v, %v\n", t.s, t.t)
		}
	}
}
