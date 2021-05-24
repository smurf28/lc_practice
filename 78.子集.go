package leetcode

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
var ret [][]int

func subsets(nums []int) [][]int {
	ret = [][]int{}
	reverse(nums, 0, []int{})
	return ret
}

// 全排列
func reverse(nums []int, i int, sub []int) {
	// fmt.Printf("format string %d,%v\n", i, sub)
	track := make([]int, len(sub))
	copy(track, sub)
	ret = append(ret, track)
	if i == len(nums) {
		return
	}

	for j := i; j < len(nums); j++ {
		sub = append(sub, nums[j])
		reverse(nums, j+1, sub)
		sub = sub[:len(sub)-1]
	}
}

// @lc code=end
