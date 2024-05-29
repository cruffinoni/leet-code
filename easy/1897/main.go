package main

import "log"

func makeEqual(words []string) bool {
	switch len(words) {
	case 0:
		return false
	case 1:
		return true
	}
	l := len(words)
	m := make(map[rune]int, len(words))
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			m[rune(words[i][j])]++
		}
	}
	for _, i := range m {
		if i%l != 0 {
			return false
		}
	}
	return true
}

func main() {
	log.Printf("makeEqual: %v/true", makeEqual([]string{"abc", "aabc", "bc"}))
}
