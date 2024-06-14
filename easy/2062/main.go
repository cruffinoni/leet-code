package main

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func countVowelSubstrings(word string) int {
	l := len(word)
	if l == 0 {
		return 0
	}
	totalCount := 0
	for i := 0; i < l; i++ {
		if isVowel(word[i]) {
			m := make(map[byte]int, 5)
			k := i
			for k < l && isVowel(word[k]) {
				m[word[k]]++
				if len(m) == 5 {
					totalCount++
				}
				k++
			}
		}
	}
	return totalCount
}

type test struct {
	word     string
	expected int
}

func main() {
	tests := []test{
		{word: "aeiouu", expected: 2},
		{word: "unicornarihan", expected: 0},
		{word: "cuaieuouac", expected: 7},
	}
	for _, t := range tests {
		result := countVowelSubstrings(t.word)
		if result != t.expected {
			println("Error: expected", t.expected, "got", result)
		} else {
			println("Success: expected", t.expected, "got", result)
		}
	}
}
