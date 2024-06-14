package main

import (
	"log"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a

}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	res := 0
	for i := 0; i < len(seats); i++ {
		res += abs(seats[i] - students[i])
	}
	return res
}

type test struct {
	seats    []int
	students []int
	expected int
}

func main() {
	tests := []test{
		{seats: []int{3, 1, 5}, students: []int{2, 7, 4}, expected: 4},
		{seats: []int{4, 1, 5, 9}, students: []int{1, 3, 2, 6}, expected: 7},
		{seats: []int{2, 2, 6, 6}, students: []int{1, 3, 2, 6}, expected: 4},
	}
	for _, t := range tests {
		res := minMovesToSeat(t.seats, t.students)
		if res != t.expected {
			log.Printf("Test failed for %v, %v, expected %d, got %v", t.seats, t.students, t.expected, res)
		} else {
			log.Printf("Test passed for %v, %v", t.seats, t.students)
		}
	}
}
