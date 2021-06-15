package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i, v := range nums {
		j, ok := lookup[-v]
		lookup[v-target] = i
		if ok {
			return []int{j, i}
		}
	}
	return []int{}
}

// @lc code=end
