package main

import (
	"log"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	l := len(dictionary)
	if l == 0 {
		return sentence
	}
	m := make(map[string]int)
	for i := 0; i < l; i++ {
		m[dictionary[i]] = len(dictionary[i])
	}
	words := strings.Split(sentence, " ")
	for i := range words {
		for d := range m {
			if strings.HasPrefix(words[i], d) {
				words[i] = d
			}
		}
	}
	return strings.Join(words, " ")
}

type test struct {
	dictionary []string
	sentence   string
	expected   string
}

func main() {
	tests := []test{
		{dictionary: []string{"cat", "bat", "rat"}, sentence: "the cattle was rattled by the battery", expected: "the cat was rat by the bat"},
		{dictionary: []string{"a", "b", "c"}, sentence: "aadsfasf absbs bbab cadsfafs", expected: "a a b c"},
		{dictionary: []string{"a", "aa", "aaa", "aaaa"}, sentence: "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa", expected: "a a a a a a a a bbb baba a"},
		{dictionary: []string{"catt", "cat", "bat", "rat"}, sentence: "the cattle was rattled by the battery", expected: "the cat was rat by the bat"},
	}
	for _, t := range tests {
		res := replaceWords(t.dictionary, t.sentence)
		if res != t.expected {
			log.Printf("Test failed for '%v', expected '%v', got '%v'", t.sentence, t.expected, res)
		} else {
			log.Printf("Test passed for %v", t.sentence)
		}
	}
}
