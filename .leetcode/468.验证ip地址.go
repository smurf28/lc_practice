/*
 * @lc app=leetcode.cn id=468 lang=golang
 *
 * [468] 验证IP地址
 */
package leetcode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// @lc code=start
func validIPAddress(IP string) string {
	if validIPv4Address(IP) {
		return "IPv4"
	}
	if validIPv6Address(IP) {
		return "IPv6"
	}
	return "Neither"
}

func validIPv4Address(IP string) bool {
	strArr := strings.Split(IP, ".")
	if len(strArr) != 4 {
		return false
	}
	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil || num > 255 || num < 0 { //注意：err != nil
			return false
		}
		newStr := fmt.Sprint(num)
		if str != newStr {
			return false
		}
	}
	return true
}

func validIPv6Address(IP string) bool {
	strArr := strings.Split(IP, ":")
	if len(strArr) != 8 {
		return false
	}
	re := regexp.MustCompile(`^([0-9]|[a-f]|[A-F])+$`)
	for _, str := range strArr {
		if len(str) == 0 || len(str) > 4 {
			return false
		}
		if !re.MatchString(str) {
			return false
		}
	}
	return true
}

// @lc code=end
