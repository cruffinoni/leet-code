package main

import (
	"log"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, maxVal *float64) int {
	if node == nil {
		return 0
	}
	leftPath := max(dfs(node.Left, maxVal), 0)
	rightPath := max(dfs(node.Right, maxVal), 0)

	total := node.Val + leftPath + rightPath
	*maxVal = max(*maxVal, float64(total))
	return node.Val + max(leftPath, rightPath)
}

func maxPathSum(root *TreeNode) int {
	m := math.Inf(-1)
	dfs(root, &m)
	log.Printf("Res: %f", m)
	return int(m)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	log.Printf("Res: %v", maxPathSum(root))
}
