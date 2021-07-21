package bytedance

import (
	"fmt"
	"lc/common"
	"testing"
)

// 翻转数组 中位数
func rotated(nums []int) int {
	l, r := 0, len(nums)-1
	n := len(nums)
	provi := -1
	for r >= l {
		mid := (r + l) / 2
		if nums[0] <= nums[mid-1] && nums[mid] <= nums[len(nums)-1] {
			provi = mid
			break
		} else if nums[0] <= nums[mid-1] {
			l = mid + 1

		} else if nums[mid] <= nums[len(nums)-1] {
			r = mid - 1
		}
	}
	if n%2 == 1 {
		return nums[(provi+n/2)%n]
	} else {
		// fmt.Print("111111111")
		return (nums[(provi+n/2)%n] + nums[(provi+n/2-1)%n])
	}
}

func TestRotated(t *testing.T) {
	a := rotated([]int{4, 5, 6, 1, 2, 3})
	fmt.Printf("%d\n", a)
}

// 二叉树最远两个节点

func maxDepth(tree *common.TreeNode) int {
	if tree == nil {
		return 0
	}
	L := maxDepth(tree.Left)
	R := maxDepth(tree.Right)
	return max(L, R) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func treeDistance(head *common.TreeNode) int {
	ld := maxDepth(head.Left)
	rd := maxDepth(head.Right)
	return ld + rd
}
func TestMaxDepth(t *testing.T) {
	head := &common.TreeNode{1, &common.TreeNode{2, &common.TreeNode{3, &common.TreeNode{5, nil, nil}, nil}, nil}, &common.TreeNode{4, nil, nil}}
	a := maxDepth(head)
	fmt.Println(a)
}
