package main

import "sort"

func heightChecker(heights []int) int {
	cpy := make([]int, len(heights))
	copy(cpy, heights)
	sort.Ints(cpy)
	count := 0
	for i := 0; i < len(heights); i++ {
		if cpy[i] != heights[i] {
			count++
		}
	}
	return count
}

func main() {

}
