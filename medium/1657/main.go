package main

import (
	"maps"
)

type data struct {
	occurrences int
}

func closeStrings(word1 string, word2 string) bool {
	l := len(word1)
	l2 := len(word2)
	if l != l2 {
		return false
	}
	occurrences1 := make(map[uint8]int)
	occurrences2 := make(map[uint8]int)
	countMap1 := make(map[int]int)
	countMap2 := make(map[int]int)

	for i := 0; i < l; i++ {
		occurrences1[word1[i]]++
		occurrences2[word2[i]]++
		countMap1[occurrences1[word1[i]]]++
		countMap2[occurrences2[word2[i]]]++
	}
	for char := range occurrences1 {
		if _, ok := occurrences2[char]; !ok {
			return false
		}
	}
	if maps.Equal(occurrences1, occurrences2) {
		return true
	}
	if !maps.Equal(countMap1, countMap2) {
		return false
	}
	return true
}

type test struct {
	word1    string
	word2    string
	expected bool
}

func main() {
	tests := []test{
		{word1: "abc", word2: "bca", expected: true},
		{word1: "a", word2: "aa", expected: false},
		{word1: "cabbba", word2: "abbccc", expected: true},
		{word1: "uau", word2: "ssx", expected: false},
	}
	for _, t := range tests {
		result := closeStrings(t.word1, t.word2)
		if result != t.expected {
			println("Error: expected", t.expected, "got", result)
		} else {
			println("Success: expected", t.expected, "got", result)
		}
	}
}
