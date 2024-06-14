package main

import "log"

func nextGreatestLetter(letters []byte, target byte) byte {
	l := len(letters)
	if target == 'z' || target < letters[0] || target > letters[l-1] {
		return letters[0]
	}
	lo, hi := 0, l-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if letters[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return letters[lo%l]
}

type test struct {
	letters  []byte
	target   byte
	expected byte
}

func main() {
	tests := []test{
		{letters: []byte{'c', 'f', 'j'}, target: 'a', expected: 'c'},
		{letters: []byte{'c', 'f', 'j'}, target: 'c', expected: 'f'},
		{letters: []byte{'x', 'x', 'y', 'y'}, target: 'z', expected: 'x'},
		{letters: []byte{'c', 'f', 'j'}, target: 'd', expected: 'f'},
		{letters: []byte{'e', 'e', 'g', 'g'}, target: 'g', expected: 'e'},
	}
	for _, t := range tests {
		res := nextGreatestLetter(t.letters, t.target)
		if res != t.expected {
			log.Printf("Test failed for %v, expected %c, got %c", t.letters, t.expected, res)
		} else {
			log.Printf("Test passed for %v", t.letters)
		}
	}
}
