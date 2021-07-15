/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */
package leetcode

import "fmt"

// @lc code=start
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		fmt.Print(l, " ", r, "\n")
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid+1
			continue
		} else {
			r = mid-1
			continue
		}
	}
	return -1
}

// @lc code=end
