package main

import "runtime/debug"

func removeStars(s string) string {
	l := len(s)
	lStack := 0
	stack := make([]byte, 0, len(s))
	for i := 0; i < l; i++ {
		if s[i] != '*' {
			stack = append(stack, s[i])
			lStack++
		} else {
			stack = stack[:lStack-1]
			lStack--
		}
	}
	return string(stack)
}

func init() {
	debug.SetGCPercent(10)
}

type test struct {
	s        string
	expected string
}

func main() {
	tests := []test{
		{s: "leet**cod*e", expected: "lecoe"},
		{s: "erase*****", expected: ""},
	}
	for _, t := range tests {
		result := removeStars(t.s)
		if result != t.expected {
			println("Error: expected", t.expected, "got", result)
		} else {
			println("Success: expected", t.expected, "got", result)
		}
	}
}
