/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
package leetcode

import (
	"math"
)

// @lc code=start

func coinChange(coins []int, amount int) int {
	// sort.Slice(coins, func(i, j int) bool { return i > j })
	if amount < 1 && len(coins) < 1 {
		return -1
	}
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			if coin <= i && dp[i] > dp[i-coin]+1 {
				dp[i] = dp[i-coin] + 1
				// fmt.Println(i, "", coin)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func coinChange1(coins []int, amount int) int {
	if amount < 1 && len(coins) < 1 {
		return -1
	}
	memo := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		memo[i] = math.MaxInt32
		for _, c := range coins {
			if i >= c && memo[i] > memo[i-c]+1 {
				memo[i] = memo[i-c] + 1
			}
		}
	}
	if memo[amount] == math.MaxInt32 {
		return -1
	}
	return memo[amount]

}

// @lc code=end
