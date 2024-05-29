package main

import (
	"log"
)

// Element represents a value in the array and the steps required to ensure it contributes to a non-decreasing array
type Element struct {
	value int
	steps int
}

// totalSteps calculates the number of steps to make the array non-decreasing
func totalSteps(nums []int) int {
	stack := []Element{}
	maxSteps := 0

	for _, num := range nums {
		currentSteps := 0
		// Remove elements from the stack that are less than or equal to the current element
		for len(stack) > 0 && stack[len(stack)-1].value <= num {
			currentSteps = max(currentSteps, stack[len(stack)-1].steps)
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			currentSteps++
		} else {
			currentSteps = 0
		}
		maxSteps = max(maxSteps, currentSteps)
		stack = append(stack, Element{value: num, steps: currentSteps})
	}

	return maxSteps
}

func main() {
	tests := map[int][]int{
		3:  {5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11},
		0:  {4, 5, 7, 7, 13},
		6:  {10, 1, 2, 3, 4, 5, 6, 1, 2, 3},
		18: {82, 228, 317, 559, 670, 1535, 1977, 890, 1243, 1502, 1720, 1808, 1819, 1966, 105, 170, 194, 320, 385, 433, 607, 633, 1144, 1195, 1365, 1490, 1778, 921, 1560, 6, 68, 69, 100, 341, 595, 679, 725, 775, 887, 1316, 296, 613, 658, 682, 777, 1203, 1350, 1431, 161, 374, 784, 794, 863, 1080, 1149, 1509, 1525, 128, 437, 996, 1045, 1061, 1102, 1238, 1624, 1706, 1961, 808, 950, 1166, 1531, 1537, 1732, 866, 1279, 1494, 1527, 1595},
	}

	for expected, input := range tests {
		cpyInput := make([]int, len(input))
		copy(cpyInput, input)
		res := totalSteps(input)
		if res != expected {
			log.Printf("expected %d, got %d with input %v", expected, res, cpyInput)
		} else {
			log.Printf("test with input %v passed!", cpyInput)
		}
	}
}
