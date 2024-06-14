package main

func oneBit(a, i int) {
	return
}

func minFlips(a int, b int, c int) int {
	flipCount := 0
	m := min(31, max(a, b, c))
	for i := 0; i < m; i++ {
		a := a & (1 << i)
		b := b & (1 << i)
		c := c & (1 << i)
		if a|b != c {
			if c == 0 && a == b && a > 0 {
				flipCount += 2
			} else {
				flipCount++
			}
		}
	}
	return flipCount
}

type test struct {
	a, b, c, expected int
}

func main() {
	tests := []test{
		{a: 2, b: 6, c: 5, expected: 3},
		{a: 4, b: 2, c: 7, expected: 1},
		{a: 1, b: 2, c: 3, expected: 0},
		{a: 8, b: 3, c: 5, expected: 3},
		{a: 10, b: 9, c: 1, expected: 3},
		{a: 851603914, b: 632925160, c: 574207408, expected: 17},
	}
	for _, t := range tests {
		result := minFlips(t.a, t.b, t.c)
		if result != t.expected {
			println("Error: expected", t.expected, "got", result)
		} else {
			println("Success: expected", t.expected, "got", result)
		}
	}
}
