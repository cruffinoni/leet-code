package main

func longestPalindrome(s string) int {
	m := make(map[uint8]int)
	for i := range s {
		m[s[i]]++
	}
	count := 0
	oddOccurrence := 0
	for _, i := range m {
		if i%2 == 0 {
			count += i
		} else {
			count += i - 1
			oddOccurrence = 1
		}
	}
	return count + oddOccurrence
}

type test struct {
	s        string
	expected int
}

/*
bananas
b=1
a=3
n=2

*/

func main() {
	tests := []test{
		{s: "abccccdd", expected: 7},
		{s: "a", expected: 1},
		{s: "bb", expected: 2},
		{s: "ccd", expected: 3},
		{s: "bananas", expected: 5},
	}
	for _, t := range tests {
		result := longestPalindrome(t.s)
		if result != t.expected {
			println("Error: expected", t.expected, "got", result)
		} else {
			println("Success: expected", t.expected, "got", result)
		}
	}
}
