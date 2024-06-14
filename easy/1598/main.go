package main

func minOperations(logs []string) int {
	//stack := make([]string, 0)
	size := 0
	for _, l := range logs {
		switch l {
		case "../":
			if size > 0 {
				size--
			}
		case "./":
			continue
		default:
			size++
		}
	}
	return size
}

type test struct {
	logs     []string
	expected int
}

func main() {
	tests := []test{
		{logs: []string{"d1/", "d2/", "../", "d21/", "./"}, expected: 2},
		{logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}, expected: 3},
		{logs: []string{"d1/", "../", "../", "../"}, expected: 0},
	}
	for _, t := range tests {
		result := minOperations(t.logs)
		if result != t.expected {
			println("Error: expected", t.expected, "got", result)
		} else {
			println("Success: expected", t.expected, "got", result)
		}
	}
}
