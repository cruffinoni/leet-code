package main

import (
	"log"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	woSpaces := strings.Fields(s)
	l := len(woSpaces)
	hi := l - 1
	lo := 0
	for hi > lo {
		woSpaces[hi], woSpaces[lo] = woSpaces[lo], woSpaces[hi]
		lo++
		hi--
	}
	return strings.Join(woSpaces, " ")
}

func main() {
	tests := map[string]string{
		"the sky is blue":  "blue is sky the",
		"  hello world  ":  "world hello",
		"a good   example": "example good a",
	}

	for input, expected := range tests {
		res := reverseWords(input)
		if res != expected {
			log.Printf("expected '%v', got '%v' for '%s'", expected, res, input)
		} else {
			log.Printf("test %v passed!", input)
		}
	}
}
