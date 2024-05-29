package main

import "log"

// https://leetcode.com/problems/subsets/description/?envType=daily-question&envId=2024-05-21

func subsets_bin(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	l := len(nums)
	mask := 1 << l
	res := make([][]int, 0, len(nums))
	for i := 0; i < mask; i++ {
		subRes := make([]int, 0, l)
		for j, k := range nums {
			if (1<<j)&i > 0 {
				subRes = append(subRes, k)
			}
		}
		res = append(res, subRes)
	}
	return res
}

// Fonction de backtracking pour générer tous les sous-ensembles
func backtrack(nums []int, index int, currentSubset []int, result *[][]int) {
	// Ajouter une copie du sous-ensemble courant au résultat
	subsetCopy := make([]int, len(currentSubset))
	copy(subsetCopy, currentSubset)
	*result = append(*result, subsetCopy)

	// Explorer toutes les possibilités à partir de l'index courant
	for i := index; i < len(nums); i++ {
		// Inclure l'élément courant dans le sous-ensemble
		currentSubset = append(currentSubset, nums[i])
		// Appeler récursivement pour inclure le prochain élément
		backtrack(nums, i+1, currentSubset, result)
		// Exclure l'élément courant pour backtrack
		currentSubset = currentSubset[:len(currentSubset)-1]
	}
}

// Fonction pour générer tous les sous-ensembles d'un tableau
func subsets_recursive(nums []int) [][]int {
	result := [][]int{}
	currentSubset := []int{}
	backtrack(nums, 0, currentSubset, &result)
	return result
}

func main() {
	tests := [][]int{
		{0},
		{1, 2, 3},
	}
	for _, i := range tests {
		log.Printf("Input %+v > %+v", i, subsets_recursive(i))
	}
}
