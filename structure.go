package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type BTree struct {
	Val    int
	LChild *BTree
	RChild *BTree
}

// 遍历
// 前序遍历
func preorderTraversal(root *BTree) {
	if root == nil {
		return
	}
	fmt.Printf("%d", root.Val)
	preorderTraversal(root.LChild)
	preorderTraversal(root.RChild)
}

func preorderTraversal1(root *BTree) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*BTree, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.LChild
		}
		// pop
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = top.RChild
	}
	return result
}

func inorderTraversal(root *BTree) {
	if root == nil {
		return
	}

	preorderTraversal(root.LChild)
	fmt.Printf("%d", root.Val)
	preorderTraversal(root.RChild)
}

func inorderTraversal1(root *BTree) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*BTree, 0)

	for root != nil || len(stack) != 0 {
		// left
		for root != nil {
			stack = append(stack, root)
			root = root.LChild
		}
		if len(stack) != 0 {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// print
			// fmt.Printf("%d", top.Val)
			result = append(result, top.Val)
			// right
			root = top.RChild
		}

	}
	return result
}

func postorderTraversal(root *BTree) {
	if root == nil {
		return
	}

	preorderTraversal(root.LChild)
	preorderTraversal(root.RChild)
	fmt.Printf("%d", root.Val)
}
