/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package leetcode

// @lc code=start
// 滑动窗口
func lengthOfLongestSubstring1(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	left, right := 0, 0
	// key：字符 value:出现位置
	visited := make(map[byte]int)
	visited[s[0]] = left
	maxLength, length := 0, 1

	for right+1 < len(s) {
		// fmt.Print(length, left, " ", s[right], " ")
		right++
		if index, ok := visited[s[right]]; ok {
			if maxLength < length {
				maxLength = length
			}
			left = max(left, index+1)
			length = right - left + 1
		} else {
			length++
		}
		visited[s[right]] = right
	}
	if maxLength < length {
		maxLength = length
	}

	return maxLength
}

func lengthOfLongestSubstring(s string) int {
	freq := make([]int, 128)
	var res = 0
	start, end := 0, -1
	for start < len(s) {
		if end+1 < len(s) && freq[s[end+1]] == 0 {
			end++
			freq[s[end]]++
		} else {
			freq[s[start]]--
			start++
		}
		res = max(res, end-start+1)
	}
	return res
}
// func max(i, j int) int {
// 	if i > j {
// 		return i
// 	} else {
// 		return j
// 	}
// }

// 滑动窗口
// func lengthOfLongestSubstring(string s) {
//     unordered_map<char, int> window;

//     int left = 0, right = 0;
//     int res = 0; // 记录结果
//     while (right < s.size()) {
//         char c = s[right];
//         right++;
//         // 进行窗口内数据的一系列更新
//         window[c]++;
//         // 判断左侧窗口是否要收缩
//         while (window[c] > 1) {
//             char d = s[left];
//             left++;
//             // 进行窗口内数据的一系列更新
//             window[d]--;
//         }
//         // 在这里更新答案
//         res = max(res, right - left);
//     }
//     return res;
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// @lc code=end
