package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start

var numberMap = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

var res []string

func letterCombinations(digits string) []string {
	res = []string{}
	if digits == "" {
		return []string{}
	}

	var digitStr []string
	for i, _ := range digits {
		digitStr = append(digitStr, string(digits[i]))
	}

	combinations(0, digitStr, []string{})
	return res
}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

// 组合
func combinations(index int, digitStr []string, str []string) {
	if index == len(digitStr) {
		res = append(res, strings.Join(str, ""))
		return
	}

	for j := 0; j < len(numberMap[digitStr[index]]); j++ {
		str = append(str, numberMap[digitStr[index]][j])
		combinations(index+1, digitStr, str)
		str = remove(str, len(str)-1)
	}
}

// @lc code=end
