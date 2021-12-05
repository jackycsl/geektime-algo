package main

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
func buildTree(inorder []int, postorder []int) *TreeNode {

	var build func(int, int) *TreeNode

	build = func(in_left int, in_right int) *TreeNode {
		if in_left > in_right {
			return nil
		}
		// fmt.Println(postorder, in_left, in_right)
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{val, nil, nil}

		mid := in_left
		for inorder[mid] != val {
			mid++
		}

		root.Right = build(mid+1, in_right)
		root.Left = build(in_left, mid-1)

		return root
	}

	return build(0, len(inorder)-1)
}
