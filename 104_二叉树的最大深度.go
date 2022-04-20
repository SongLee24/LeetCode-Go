package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	dep := depth(root, 0)
	return dep
}

// 递归
func depth(root *TreeNode, dep int) int {
	if root == nil {
		return dep
	}

	leftDep := depth(root.Left, dep+1)
	rightDep := depth(root.Right, dep+1)

	if leftDep > rightDep {
		return leftDep
	} else {
		return rightDep
	}
}
