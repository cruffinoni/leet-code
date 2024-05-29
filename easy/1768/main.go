package main

import "log"

func mergeAlternately(word1 string, word2 string) string {
	la := len(word1)
	lb := len(word2)
	if la == 0 {
		return word2
	} else if lb == 0 {
		return word1
	}
	res := make([]byte, 0, la+lb)
	for i := range la + lb {
		if i < la {
			res = append(res, word1[i])
		}
		if i < lb {
			res = append(res, word2[i])
		}
	}
	return string(res)
}

type test struct {
	word1 string
	word2 string
}

func main() {
	tests := map[string]test{
		"apbqcr": {word1: "abc", word2: "pqr"},
		"apbqrs": {word1: "ab", word2: "pqrs"},
		"apbqcd": {word1: "abcd", word2: "pq"},
	}
	for expected, t := range tests {
		res := mergeAlternately(t.word1, t.word2)
		if res != expected {
			log.Printf("expected '%s', got '%s'", expected, res)
		} else {
			log.Printf("test '%s' passed!", expected)
		}
	}
}
