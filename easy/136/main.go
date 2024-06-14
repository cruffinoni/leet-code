package main

func singleNumber(nums []int) int {
	nb := 0
	for i := 0; i < len(nums); i++ {
		nb ^= nums[i]
	}
	return nb
}

type test struct {
	nums     []int
	expected int
}

func main() {
	tests := []test{
		{nums: []int{2, 2, 3}, expected: 3},
		{nums: []int{0, 1, 0, 1, 0, 1, 99}, expected: 99},
		{nums: []int{1}, expected: 1},
	}
	for _, t := range tests {
		result := singleNumber(t.nums)
		if result != t.expected {
			println("Error: expected", t.expected, "got", result)
		} else {
			println("Success: expected", t.expected, "got", result)
		}
	}
}
