/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */
package leetcode

import "fmt"

// @lc code=start

func uniquePaths(m int, n int) int {
	// dp := [m][n]int{}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[1][1] = 1

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}

func reversePaths(col, row, m, n int) int {
	if col == m && row == n {
		fmt.Printf("row %d col %d\n", col, row)
		return 1
	}
	if col > m || row > n {
		return 0
	}
	if col < m || row < n {
		return reversePaths(col+1, row, m, n) + reversePaths(col, row+1, m, n)
	}
	return 0
}

// @lc code=end
