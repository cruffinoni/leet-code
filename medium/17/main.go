package main

import (
	"log"
	"reflect"
	"strings"
)

var digitToLetters = map[byte]string{
	'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
	'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
}

func backtrack(res *[]string, digits string, digitsLen, currentIdx int, combination string) {
	if len(combination) == digitsLen {
		*res = append(*res, combination)
		return
	}
	currentLetters := digitToLetters[digits[currentIdx]]
	for _, b := range currentLetters {
		backtrack(res, digits, digitsLen, currentIdx+1, combination+string(b))
	}
}

func letterCombinations(digits string) []string {
	digits = strings.ReplaceAll(digits, "1", "")
	if digits == "" {
		return []string{}
	}
	l := len(digits)
	res := make([]string, 0)
	backtrack(&res, digits, l, 0, "")
	return res
}

type test struct {
	digits   string
	expected []string
}

func main() {
	tests := []test{
		{digits: "23", expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{digits: "", expected: []string{}},
		{digits: "2", expected: []string{"a", "b", "c"}},
	}
	for _, t := range tests {
		res := letterCombinations(t.digits)
		if !reflect.DeepEqual(res, t.expected) {
			log.Printf("Test failed for %s, expected %v, got %v", t.digits, t.expected, res)
		} else {
			log.Printf("Test passed for %s", t.digits)
		}
	}
}
