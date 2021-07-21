/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
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
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid, end := head, head
	for {
		if end.Next == nil || (end.Next != nil && end.Next.Next == nil) {
			break
		}
		end = end.Next.Next
		mid = mid.Next
	}

	toMerge := reverseList(mid.Next)
	mid.Next = nil

	mergeList(toMerge, head)
}

func mergeList(toMerge, head *ListNode) {

	for toMerge != nil && head != nil {
		headTemp := head.Next
		toMergeTemp := toMerge.Next

		head.Next = toMerge
		toMerge.Next = headTemp

		toMerge = toMergeTemp
		head = headTemp
	}
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

// @lc code=end
