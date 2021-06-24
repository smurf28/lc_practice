package leetcode

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 */
type NodeLevel struct {
	node *TreeNode
	// 层数
	order int
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	var res [][]int
	// 入队
	queue = append(queue, root)

	for len(queue) != 0 {
		count := len(queue)
		levelNums := []int{}
		for count > 0 {
			topNode := queue[0]
			levelNums = append(levelNums, queue[0].Val)
			// pop queue
			queue = queue[1:]
			if topNode.Left != nil {
				queue = append(queue, topNode.Left)
			}

			if topNode.Right != nil {
				queue = append(queue, topNode.Right)
			}
			count--
		}
		res = append(res, levelNums)
	}
	return res
}

// @lc code=end
