package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFarmLand(t *testing.T) {
	cases := []struct {
		land     [][]int
		expected [][]int
	}{
		{
			land: [][]int{
				{1, 0, 0},
				{0, 1, 1},
				{0, 1, 1},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{1, 1, 2, 2},
			},
		},
	}
	for i, c := range cases {
		t.Run("Case_"+strconv.Itoa(i), func(t *testing.T) {
			res := findFarmland(c.land)
			assert.Equal(t, c.expected, res)
		})
	}
}
