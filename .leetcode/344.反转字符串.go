package leetcode

/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start

// 递归解法
func reverseString(s []byte) {
	reverseS(s, 0)
}

func reverseS(s []byte, index int) {
	right := len(s) - index - 1
	if index >= right {
		return
	}

	s[index] ^= s[right]
	s[right] ^= s[index]
	s[index] ^= s[right]
	reverseS(s, index+1)
}

// @lc code=end
