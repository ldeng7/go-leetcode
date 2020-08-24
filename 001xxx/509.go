import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minDifference(nums []int) int {
	l := len(nums)
	if l <= 4 {
		return 0
	}
	sort.Ints(nums)
	return min(min(nums[l-4]-nums[0], nums[l-3]-nums[1]), min(nums[l-2]-nums[2], nums[l-1]-nums[3]))
}
