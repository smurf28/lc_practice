package leetcode

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
// 双向指针
func maxArea(height []int) int {
	left_i := 0
	right_i := len(height) - 1
	res := 0
	for {
		if left_i >= right_i {
			break
		}
		res = max(res, (right_i-left_i)*min(height[right_i], height[left_i]))
		if height[right_i] < height[left_i] {
			right_i--
		} else {
			left_i++
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
