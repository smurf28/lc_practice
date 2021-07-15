/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package leetcode

// @lc code=start
func longestPalindrome(s string) string {
	n := len(s)
	var res string
	dp := make([]bool, n)

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= i; j-- {
			dp[j] = (s[i] == s[j]) && (j-i < 3 || dp[j-1])
			if dp[j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}

	return res

}

// @lc code=end
