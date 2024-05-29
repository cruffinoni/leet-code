package main

import "log"

func lengthOfLongestSubstring(s string) int {
	charSeen := map[uint8]int{}
	maxLen := 0
	currLen := 0
	for i := 0; i < len(s); i++ {
		if _, ok := charSeen[s[i]]; !ok {
			currLen++
			charSeen[s[i]] = i
		} else {
			maxLen = max(currLen, maxLen)
			currLen = 1
			i = charSeen[s[i]] + 1
			charSeen = map[uint8]int{
				s[i]: i,
			}
		}
	}
	return max(currLen, maxLen)
}

func main() {
	tests := map[string]int{
		"abcabcbb":  3,
		"bbbbb":     1,
		"pwwkew":    3,
		"aab":       2,
		"dvdf":      3,
		"bpfbhmipx": 7,
	}
	for input, expected := range tests {
		res := lengthOfLongestSubstring(input)
		if res != expected {
			log.Printf("expected %d, got %d for %s", expected, res, input)
		} else {
			log.Printf("test %v passed!", input)
		}
	}
}
