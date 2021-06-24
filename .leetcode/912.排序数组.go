package leetcode

/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start

func sortArray(nums []int) []int {
	// quikSort(nums, 0, len(nums)-1)
	heapSort(nums)
	return nums
}

// 快速排序
func quikSort(nums []int, start, end int) {
	if end > start {
		privotpos := partition(nums, start, end)
		// 递归 排序左边
		quikSort(nums, start, privotpos-1)
		quikSort(nums, privotpos+1, end)
	}
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	// 小于pivote的放左边，大于pivote 放右边
	for start != end {
		for nums[end] >= pivot && start < end {
			end--
		}
		nums[start] = nums[end]

		for nums[start] <= pivot && start < end {
			start++
		}
		nums[end] = nums[start]
	}
	nums[start] = pivot
	return start
}
func partition1(nums []int, start, end int) int {
	pos := nums[start]
	for start != end {
		if nums[end] >= pos && start < end {
			end--
		}
		nums[start] = nums[end]
		if nums[start] <= pos && start < end {
			start++
		}
		nums[end] = nums[start]
	}
	nums[start] = pos
	return start
}

// 堆排序
func heapSort(nums []int) {
	// 构建大顶堆
	// len(nums)/2 - 1 为最后一个非叶子节点
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjust(nums, len(nums), i)
	}

	for i := len(nums) - 1; i > 0; i-- {
		swap(&nums[0], &nums[i])
		adjust(nums, i, 0)
	}
}

func adjust(nums []int, len, index int) {
	left := 2*index + 1
	right := 2*index + 2
	maxIdx := index

	if left < len && nums[left] > nums[maxIdx] {
		maxIdx = left
	}
	if right < len && nums[right] > nums[maxIdx] {
		maxIdx = right
	}

	if maxIdx != index {
		swap(&nums[maxIdx], &nums[index])
		adjust(nums, len, maxIdx)
	}
}

func swap(a, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}

// 归并排序
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	// 分治法：divide 分为两段
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	// 合并两段数据
	result := merge(left, right)
	return result
}
func merge(left, right []int) (result []int) {
	// 两边数组合并游标
	l := 0
	r := 0
	// 注意不能越界
	for l < len(left) && r < len(right) {
		// 谁小合并谁
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	// 剩余部分合并
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

// 冒泡排序
func Sort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	swaped := true
	j := 0
	for swaped {
		swaped = false
		for i := 1; i < len(nums)-j; i++ {
			if nums[i-1] > nums[i] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				swaped = true
			}
		}
		j++
	}
}

// @lc code=end
