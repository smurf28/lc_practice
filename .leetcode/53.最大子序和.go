/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
package leetcode

import "math"

// @lc code=start
func maxSubArray(nums []int) int {
	res, sum := math.MinInt32, -1
	for i := 0; i < len(nums); i++ {
		sum = max(nums[i], sum+nums[i])
		res = max(res, sum)
	}
	return res
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// @lc code=end
