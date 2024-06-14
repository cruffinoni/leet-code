package main

import "log"

func commonChars(words []string) []string {
	length := len(words)
	if len(words) == 1 {
		return words
	}
	m := make([]map[byte]int, len(words))
	for i := 0; i < length; i++ {
		m[i] = make(map[byte]int)
		for c := range words[i] {
			m[i][words[i][c]]++
		}
	}
	end := make([]string, 0)
	for k, v := range m[0] {
		countFound := 1
		for i := 1; i < length; i++ {
			if m[i][k] == 0 {
				break
			}
			countFound++
			v = min(v, m[i][k])
		}
		if countFound == length {
			for i := 0; i < v; i++ {
				end = append(end, string(k))
			}
		}
	}
	return end
}

type test struct {
	words    []string
	expected []string
}

func main() {
	tests := []test{
		{words: []string{"bella", "label", "roller"}, expected: []string{"e", "l", "l"}},
		{words: []string{"cool", "lock", "cook"}, expected: []string{"c", "o"}},
	}
	for _, t := range tests {
		res := commonChars(t.words)
		if len(res) != len(t.expected) {
			log.Printf("Test failed for %v, expected %v, got %v", t.words, t.expected, res)
		} else {
			log.Printf("Test passed for %v", t.words)
		}
	}
}
