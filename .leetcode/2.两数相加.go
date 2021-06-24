/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum_res := &ListNode{Next: nil}
	pos_res := sum_res
	sub := 0 //上位 数
	for l1 != nil && l2 != nil {
		temp_sum := l1.Val + l2.Val + sub
		sub = temp_sum / 10
		pos_res.Next = &ListNode{
			Val:  temp_sum % 10,
			Next: nil,
		}
		pos_res = pos_res.Next

		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		temp_sum := l1.Val + sub
		sub = temp_sum / 10
		pos_res.Next = &ListNode{
			Val:  temp_sum % 10,
			Next: nil,
		}
		pos_res = pos_res.Next
		l1 = l1.Next
	}
	for l2 != nil {
		temp_sum := l2.Val + sub
		sub = temp_sum / 10
		pos_res.Next = &ListNode{
			Val:  temp_sum % 10,
			Next: nil,
		}
		pos_res = pos_res.Next
		l2 = l2.Next
	}

	if sub != 0 {
		pos_res.Next = &ListNode{
			Val:  sub,
			Next: nil,
		}
	}
	return sum_res.Next
}

// @lc code=end
