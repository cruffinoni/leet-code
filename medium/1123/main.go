package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaLeaves(root *TreeNode, depth int) (int, *TreeNode) {
	if root == nil {
		log.Printf("Deepest depth: %v", depth)
		return depth, nil
	}
	leftMax, leftNode := lcaLeaves(root.Left, depth+1)
	rightMax, rightNode := lcaLeaves(root.Right, depth+1)
	if leftMax == rightMax {
		return leftMax, root
	} else if leftMax > rightMax {
		return leftMax, leftNode
	} else {
		return rightMax, rightNode
	}
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	_, node := lcaLeaves(root, 0)
	return node
}

type test struct {
	root     *TreeNode
	expected *TreeNode
}

func main() {
	tests := []test{
		// [0,1,3,null,2]
		{
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val:   2,
						Right: nil,
						Left:  nil,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: &TreeNode{
				Val: 2,
			},
		},
	}

	for _, t := range tests {
		result := lcaDeepestLeaves(t.root)
		if result.Val != t.expected.Val {
			println("Error: expected", t.expected, "got", result)
		} else {
			println("Success: expected", t.expected, "got", result)
		}
	}
}
