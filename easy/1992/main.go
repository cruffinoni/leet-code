package main

func findFarmland(land [][]int) [][]int {
	res := make([][]int, 0)
	r1, c1, r2, c2 := 0, 0, 0, 0
	for i := range land {
		for j := range land[i] {
			if land[i][j] == 1 {
				r1, c1 = i, j
			} else {
				r2, c2 = i, j
				res = append(res, []int{r1, c1, r2, c2})
			}
		}
	}
	return res
}
