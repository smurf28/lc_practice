/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */
package leetcode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	L := maxDepth(root.Left)
	R := maxDepth(root.Right)
	return max(L, R) + 1
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// @lc code=end
