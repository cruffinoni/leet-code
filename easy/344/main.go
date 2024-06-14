package main

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

type test struct {
	s        []byte
	expected []byte
}

func main() {
	tests := []test{
		{s: []byte("coaching"), expected: []byte("gnihcaoc")},
		{s: []byte("abcde"), expected: []byte("edcba")},
		{s: []byte("z"), expected: []byte("z")},
	}
	for _, t := range tests {
		cpy := make([]byte, len(t.s))
		copy(cpy, t.s)
		reverseString(t.s)
		if string(cpy) != string(t.expected) {
			println("Error: expected", string(t.expected), "got", string(t.s))
		} else {
			println("Success: expected", string(t.expected), "got", string(t.s))
		}
	}
}
