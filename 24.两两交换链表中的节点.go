package leetcode

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

//  if head == nil || head.Next == nil {
// 	return head
// }
// p := head.Next
// head.Next = swapPairs(p.Next)
// p.Next = head
// return p
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head.Next
	head.Next = swapPairs(p.Next)
	p.Next = head
	return p
}

// @lc code=end
