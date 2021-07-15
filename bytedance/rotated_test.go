package bytedance

import (
	"fmt"
	"testing"
)

func rotated(nums []int) int {
	l, r := 0, len(nums)-1
	n := len(nums)
	provi := -1
	for r >= l {
		mid := (r + l) / 2
		if nums[0] <= nums[mid-1] && nums[mid] <= nums[len(nums)-1] {
			provi = mid
		} else if nums[0] <= nums[mid-1] {
			l = mid + 1

		} else if nums[mid] <= nums[len(nums)-1] {
			r = mid - 1
		}
	}
	if n%2 == 1 {
		return nums[(provi+n/2)%n]
	} else {
		return (nums[(provi+n/2)%n] + nums[(provi+n/2-1)%n]) / 2
	}
}

func TestRotated(t *testing.T) {
	a := rotated([]int{4, 5, 1, 2, 3})
	fmt.Printf("%d\n", a)
}
