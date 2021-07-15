/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */
package leetcode

// @lc code=start
// 备忘录，消除重叠子问题
// var memo [][]int

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	// dp := [m][n]int{}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func longestCommonSubsequence(text1 string, text2 string) int {
// 	m := len(text1)
// 	n := len(text2)
// 	// 备忘录值为 -1 代表未曾计算
// 	memo = [m][n]int{}

// 	// 计算 s1[0..] 和 s2[0..] 的 lcs 长度
// 	return dp(s1, 0, s2, 0)
// }

// @lc code=end
