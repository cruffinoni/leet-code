package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func scoreOfString(s string) int {
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		sum += abs(int(s[i]) - int(s[i+1]))
	}
	return sum
}

type test struct {
	s        string
	expected int
}

func main() {
	tests := []test{
		{s: "hello", expected: 13},
		{s: "zaz", expected: 50},
	}
	for _, t := range tests {
		result := scoreOfString(t.s)
		if result != t.expected {
			println("Error: expected", t.expected, "got", result)
		} else {
			println("Success: expected", t.expected, "got", result)
		}
	}
}
