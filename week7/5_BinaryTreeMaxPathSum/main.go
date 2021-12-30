package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// O(n) time
// O(n) space
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var preOrder func(*TreeNode) int
	preOrder = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		current_sum_left := max(preOrder(node.Left), 0)
		current_sum_right := max(preOrder(node.Right), 0)

		path_sum := node.Val + current_sum_left + current_sum_right
		maxSum = max(maxSum, path_sum)

		return node.Val + max(current_sum_left, current_sum_right)
	}
	preOrder(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
