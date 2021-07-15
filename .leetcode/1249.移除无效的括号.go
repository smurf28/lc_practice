/*
 * @lc app=leetcode.cn id=1249 lang=golang
 *
 * [1249] 移除无效的括号
 */
package leetcode

import (
	"bytes"
	"fmt"
)

// @lc code=start
func minRemoveToMakeValid(s string) string {
	stack := make([]int, 0, len(s)) //存储index位置
	res := bytes.Buffer{}
	bitmap := make([]bool, len(s))
	fmt.Println("bitmap ", bitmap)
	for pos := 0; pos < len(s); pos++ {
		if s[pos] != '(' && s[pos] != ')' {
			bitmap[pos] = true
		} else if s[pos] == '(' {
			// 入栈
			stack = append(stack, pos)
		} else if len(stack) > 0 {
			// 出栈
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			bitmap[top] = true
			bitmap[pos] = true
		}
	}
	
	for i, b := range bitmap {
		if b {
			res.WriteByte(s[i])
		}
	}
	return res.String()
}

// @lc code=end
