package tx

import (
	// . "nc_tools"
	"testing"
)

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	var pre *ListNode
	cur := pHead
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

func ReverseList1(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	last := ReverseList1(pHead.Next)
	pHead.Next.Next = pHead
	pHead.Next = nil
	return last
}

func TestReverseList(t *testing.T) {
	node3 := ListNode{3, nil}
	node2 := ListNode{2, &node3}
	head := ListNode{1, &node2}
	reversed := ReverseList(&head)
	t.Logf("%v", reversed)
}
