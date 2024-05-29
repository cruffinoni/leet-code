package main

import "log"

func numSteps(s string) int {
	steps := 0
	n := len(s)

	// Convert the string to a mutable list of runes
	bin := []rune(s)

	// Iterate from the end of the string
	for n > 1 {
		if bin[n-1] == '0' {
			// Even: Divide by 2 by removing the last bit
			bin = bin[:n-1]
		} else {
			// Odd: Add 1 and handle carry
			carry := 1
			for i := n - 1; i >= 0 && carry > 0; i-- {
				if bin[i] == '1' {
					bin[i] = '0'
				} else {
					bin[i] = '1'
					carry = 0
				}
			}
			if carry == 1 {
				// If carry is still 1, we need to add a '1' at the start
				bin = append([]rune{'1'}, bin...)
			}
		}
		n = len(bin)
		steps++
	}
	return steps
}

func main() {
	tests := map[string]int{
		"1101": 6,
		"10":   1,
		"1":    0,
		"100100001010010101101000111100100111101111000111111010010001100001000100011001": 120,
	}
	for k, v := range tests {
		result := numSteps(k)
		if result != v {
			log.Printf("Test failed for %v. Got: %v, Expected: %v\n", k, result, v)
		} else {
			log.Printf("Test passed for %v\n", k)
		}
	}
}
