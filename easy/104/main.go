package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1
}

type test struct {
	root     *TreeNode
	expected int
}

func main() {
	tests := []test{
		{
			root:     &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			expected: 3,
		},
		{
			root:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			expected: 2,
		},
		{
			root:     nil,
			expected: 0,
		},
	}
	for _, t := range tests {
		result := maxDepth(t.root)
		if result != t.expected {
			log.Printf("Test failed for %v. Got: %v, Expected: %v\n", t.root, result, t.expected)
		} else {
			log.Printf("Test passed for %v\n", t.root)
		}
	}
}
