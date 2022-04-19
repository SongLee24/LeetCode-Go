package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	output := make([][]int, 0)
	if root == nil {
		return output
	}

	subLevelOrder(root, 0, &output)

	return output
}

func subLevelOrder(root *TreeNode, level int, output *[][]int) { // 参数output需要传指针类型
	if root == nil {
		return
	}

	if level >= len(*output) {
		(*output) = append(*output, make([]int, 0)) // 初始化
	}

	(*output)[level] = append((*output)[level], root.Val)

	subLevelOrder(root.Left, level+1, output) // 递归
	subLevelOrder(root.Right, level+1, output)
}
