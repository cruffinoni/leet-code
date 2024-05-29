package main

import "log"

func maxVowels(s string, k int) int {
	l := len(s)
	if k > l {
		return 0
	}
	vowels := map[uint8]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	countVowel := 0
	for i := 0; i < k; i++ {
		if _, ok := vowels[s[i]]; ok {
			countVowel++
		}
	}
	allMax := countVowel
	left, right := 0, k
	for right < l {
		if _, ok := vowels[s[left]]; ok {
			countVowel = max(countVowel-1, 0)
		}
		if _, ok := vowels[s[right]]; ok {
			countVowel++
		}
		allMax = max(allMax, countVowel)
		left++
		right++
	}
	return allMax
}

type test struct {
	s        string
	k        int
	expected int
}

func main() {
	tests := []test{
		{
			s:        "abciiidef",
			k:        3,
			expected: 3,
		},
		{
			s:        "aeiou",
			k:        2,
			expected: 2,
		},
		{
			s:        "leetcode",
			k:        3,
			expected: 2,
		},
		{
			s:        "rhythms",
			k:        4,
			expected: 0,
		},
		{
			s:        "tryhard",
			k:        4,
			expected: 1,
		},
	}
	for _, t := range tests {
		result := maxVowels(t.s, t.k)
		if result != t.expected {
			log.Printf("Test failed for %v, %v. Got: %v, Expected: %v\n", t.s, t.k, result, t.expected)
		} else {
			log.Printf("Test passed for %v, %v\n", t.s, t.k)
		}
	}
}
