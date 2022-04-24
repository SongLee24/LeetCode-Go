package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 先序(根)遍历
func flatten(root *TreeNode) {
	var head *TreeNode = &TreeNode{}
	var tmp *TreeNode = head

	// 定义函数
	var pre func(root *TreeNode)
	pre = func(root *TreeNode) {
		if root == nil {
			return
		}

		left := root.Left
		right := root.Right

		// 构造独脚树
		tmp.Right = root
		tmp.Left = nil
		tmp = tmp.Right

		// 递归
		pre(left)
		pre(right)
	}

	pre(root)
	root = head.Right
}
