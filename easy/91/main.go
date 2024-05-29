package main

import (
	"log"
	"strconv"
)

func numDecodings(s string) int {
	res := make([]rune, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			return len(res)
		}
		if i+1 < len(s) {
			tmp := string(s[i]) + string(s[i+1])
			nb, _ := strconv.Atoi(tmp)
		}
		res = append(res, rune(s[i]-'0'))
	}
	return len(res)
}

func main() {
	tests := map[string]int{
		"12":  2,
		"226": 3,
		"06":  0,
		"27":  1,
	}
	for input, expected := range tests {
		res := numDecodings(input)
		if res != expected {
			log.Printf("Tests w/ '%s', expected %d, got %d", input, expected, numDecodings(input))
		} else {
			log.Printf("Test w/ input '%s' passed!", input)
		}
	}
}
