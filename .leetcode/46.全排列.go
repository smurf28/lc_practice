/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */
package leetcode

// @lc code=start

func permute(nums []int) [][]int {
	res := [][]int{}
	part := []int{}
	reversePermute(nums, part, 0, &res)
	return res
}

func reversePermute(nums, part []int, index int, res *[][]int) {
	if index == len(nums) {
		temp := make([]int, len(part))
		copy(temp, part)
		*res = append(*res, temp)
		return
	}
	for i := index; i < len(nums); i++ {
		nums[index], nums[i] = nums[i], nums[index]
		part = append(part, nums[index])
		reversePermute(nums, part, index+1, res)
		nums[index], nums[i] = nums[i], nums[index]
		part = part[:len(part)-1]
	}
}

// @lc code=end
