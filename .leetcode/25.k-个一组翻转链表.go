/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */
package leetcode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	to := head
	for i := 0; i < k; i++ {
		// 3. 不足k个不需要反转
		if to == nil {
			return head
		}
		to = to.Next
	}
	newHead := reverseAB(head, to)
	// 1. 注意 head 反转完 变为 尾部
	// 2. 递归时从to开始 而不是to.Next
	head.Next = reverseKGroup(to, k)
	return newHead
}

func reverseAB(a, b *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := a
	for cur != b {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

// @lc code=end
