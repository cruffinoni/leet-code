package main

import (
	"log"
	"unicode"
)

func reverseVowels(s string) string {
	voyels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	for v := range voyels {
		voyels[unicode.ToUpper(v)] = struct{}{}
	}
	b := []byte(s)
	l := len(b)
	lMax := -1
	idxMax := l - 1
	lMin := -1
	idxMin := 0
	for i := 0; i < l; i++ {
		if _, ok := voyels[rune(b[idxMax])]; ok {
			lMax = idxMax
			//log.Printf("Setting lMax: %d", lMax)
		} else {
			idxMax--
		}
		if _, ok := voyels[rune(b[idxMin])]; ok {
			lMin = idxMin
			//log.Printf("Setting lMin: %d", lMin)
		} else {
			idxMin++
		}
		if idxMin > idxMax {
			return string(b)
		}
		if lMin != -1 && lMax != -1 {
			b[lMin], b[lMax] = b[lMax], b[lMin]
			//log.Printf("Min & max non-zero: %v & %v > %s", lMin, lMax, b)
			lMin, lMax = -1, -1
			idxMin++
			idxMax--
		}
	}
	return string(b)
}

func main() {
	tests := map[string]string{
		"hello":                     "holle",
		"leetcode":                  "leotcede",
		"ai":                        "ia",
		"Euston saw I was not Sue.": "euston saw I was not SuE.",
		"Marge, let's \"went.\" I await news telegram.": "Marge, let's \"went.\" i awaIt news telegram.",
	}
	for input, expected := range tests {
		cpy := input
		res := reverseVowels(cpy)
		if res != expected {
			log.Printf("expected '%v', got '%v'", expected, res)
		} else {
			log.Printf("test %v passed!", input)
		}
	}
}
