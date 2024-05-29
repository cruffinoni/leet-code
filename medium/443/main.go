package main

import "log"

func writeToChars(chars []byte, startIdx, total int, currentChar byte) int {
	write := 1
	for total > 0 {
		write++
		res := total % 10
		total /= 10
		chars[startIdx] = byte(res + '0')
		startIdx--
	}
	chars[startIdx] = currentChar
	return write
}

func getNumSize(num int) int {
	count := 0
	for num > 0 {
		count++
		num /= 10
	}
	return count
}

func compress(chars []byte) int {
	l := len(chars)
	if l == 1 {
		return 1
	}
	countWrite := 0
	currentChar := chars[0]
	count := 1
	for i := 1; i < l; i++ {
		if currentChar != chars[i] {
			if count > 1 {
				countWrite += writeToChars(chars, countWrite+getNumSize(count), count, currentChar)
			} else {
				chars[countWrite] = currentChar
				countWrite++
			}
			currentChar = chars[i]
			count = 1
		} else {
			count++
		}
	}
	if count > 1 {
		countWrite += writeToChars(chars, countWrite+getNumSize(count), count, currentChar)
		return countWrite
	} else {
		chars[countWrite] = currentChar
		return countWrite + count
	}
}

type test struct {
	input    []byte
	expected string
	length   int
}

func main() {
	tests := []test{
		{
			input:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: "a2b2c3",
			length:   6,
		},
		{
			input:    []byte{'a'},
			expected: "a",
			length:   1,
		},
		{
			input:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected: "ab12",
			length:   4,
		},
		{
			input:    []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'},
			expected: "a3b2a2",
			length:   6,
		},
		{
			input:    []byte{'a', 'a', 'a', 'a', 'a', 'b'},
			expected: "a5b",
			length:   3,
		},
		{
			input:    []byte{'a', 'a', 'a', 'a', 'b', 'a'},
			expected: "a4ba",
			length:   4,
		},
	}
	for _, t := range tests {
		res := compress(t.input)
		if string(t.input[:res]) != t.expected {
			log.Printf("expected %s, got %s", t.expected, t.input[:res])
		} else {
			log.Printf("test %s passed!", t.expected)
		}
	}
}
