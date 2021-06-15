package common

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遍历
// 前序遍历
// func preorderTraversal(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	fmt.Printf("%d", root.Val)
// 	preorderTraversal(root.Left)
// 	preorderTraversal(root.Right)
// }

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	reverse1(root, &res)
	return res
}

func reverse1(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	reverse1(root.Left, res)
	reverse1(root.Right, res)
}

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = top.Right
	}
	return result
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	inReverse(root, &res)
	return res
}

func inReverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inReverse(root.Left, res)
	*res = append(*res, root.Val)
	inReverse(root.Right, res)
}

func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		// left
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// print
			// fmt.Printf("%d", top.Val)
			result = append(result, top.Val)
			// right
			root = top.Right
		}

	}
	return result
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	postReverse(root, &res)
	return res
}

func postReverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postReverse(root.Left, res)
	postReverse(root.Right, res)
	*res = append(*res, root.Val)
}

// 当前经过节点是叶子节点。
// 当前经过节点的右子节点是上一次访问的节点。
func postorderTraversal1(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var pre *TreeNode = nil
	if root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		if root.Right == nil || root.Right == pre {

		} else {
			stack = append(stack, root.Right)
			root = root.Right
		}
	}
	return result
}

func swap(a, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}
