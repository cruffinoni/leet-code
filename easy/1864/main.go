package main

import (
	"bytes"
	"strings"
)

func maximumOddBinaryNumber(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	b := strings.Builder{}
	countOnes := bytes.Count([]byte(s), []byte{'1'})
	if countOnes == 1 {
		return strings.Repeat("0", l-1) + "1" // Add leading zeros
	} else if countOnes == l {
		return s // Fast return
	}
	for range countOnes - 1 { // countOnes - 1 for the least meaning bit to be 1 and make the number odd
		b.WriteString("1")
	}
	for range l - countOnes {
		b.WriteString("0")
	}
	b.WriteString("1")
	return b.String()
}

type test struct {
	s        string
	expected string
}

func main() {
	tests := []test{
		{s: "010", expected: "1"},
		{s: "0101", expected: "1001"},
		{s: "111", expected: "111"},
	}
	for _, t := range tests {
		result := maximumOddBinaryNumber(t.s)
		if result != t.expected {
			println("Error: expected", t.expected, "got", result)
		} else {
			println("Success: expected", t.expected, "got", result)
		}
	}
}
