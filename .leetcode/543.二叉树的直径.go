/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
// var res int = 0

// func maxDepth(tree *TreeNode) int {
// 	if tree == nil {
// 		return 0
// 	}
// 	L := maxDepth(tree.Left)
// 	R := maxDepth(tree.Right)
// 	res = max(res, L+R+1)
// 	return max(L, R) + 1
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func diameterOfBinaryTree(root *TreeNode) int {
// 	res = 1
// 	maxDepth(root)
// 	return res - 1
// }

// @lc code=end
