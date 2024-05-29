package main

import "log"

func minOperations(s string) int {
	// Variables to count the number of changes needed for both patterns
	count0, count1 := 0, 0

	for i := 0; i < len(s); i++ {
		// For pattern "0101..."
		if i%2 == 0 {
			if s[i] != '0' {
				count0++
			}
		} else {
			if s[i] != '1' {
				count0++
			}
		}
		// For pattern "1010..."
		if i%2 == 0 {
			if s[i] != '1' {
				count1++
			}
		} else {
			if s[i] != '0' {
				count1++
			}
		}
	}

	// Return the minimum of the two counts
	if count0 < count1 {
		return count0
	}
	return count1
}
func main() {
	tests := map[string]int{
		"0100":     1,
		"10":       0,
		"1111":     2,
		"1":        0,
		"10010100": 3,
	}
	for input, expected := range tests {
		res := minOperations(input)
		if expected != res {
			log.Printf("Input: '%s' expected %d, got %d", input, expected, minOperations(input))
		} else {
			log.Printf("Input %s pass!", input)
		}
	}
}

/*
10010100
- 01010101
*/
