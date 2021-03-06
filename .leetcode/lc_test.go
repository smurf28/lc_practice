package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	ret := maxArea([]int{2, 1, 1, 4})
	t.Logf("======%v", ret)
}

func TestLetterCombinations(t *testing.T) {
	ret := letterCombinations("29")
	t.Logf("ret: %v", ret)
}

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'k', 'o'}
	fmt.Println(s)
	// s = remove(s, len(s)-1)
	reverseString(s)
	fmt.Println(s)
}

func TestSwapPairs(t *testing.T) {
	head := &ListNode{Val: 1}
	head1 := &ListNode{Val: 2}
	head2 := &ListNode{Val: 3}
	head3 := &ListNode{Val: 4}
	head4 := &ListNode{Val: 5}
	head.Next = head1
	head1.Next = head2
	head2.Next = head3
	head3.Next = head4
	swapPairs(head)
	fmt.Println("head", head)
}

func TestSubsets(t *testing.T) {
	subsets([]int{1, 2, 3})
}

func TestTree(t *testing.T) {
	head := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, nil}}
	res := inorderTraversal1(head)
	fmt.Println("\ninorder", res)
	res = postorderTraversal(head)
	fmt.Println("\ninorder", res)
	// inorderTraversal(head)
}

func TestSortArrayS(t *testing.T) {
	nums := []int{2, 3, 1, 4}
	heapSort(nums)
	fmt.Println(nums)

	// a := 1
	// b := 2
	// swap(&a, &b)
	// // swap1(a, b)
	// fmt.Println("a,b:", a, b)
}

func TestLongestCommonSubsequence(t *testing.T) {
	t.Log(longestCommonSubsequence("abcde", "ace"))
}

func TestMinRemoveToMakeValid(t *testing.T) {
	t.Log(minRemoveToMakeValid("lee(t(code)"))
}

func TestReorderList(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}

	// head3.Next = head4
	reorderList(head)
	for head != nil {
		fmt.Println("head", head)
		head = head.Next
	}

	fmt.Printf("num ")

}
