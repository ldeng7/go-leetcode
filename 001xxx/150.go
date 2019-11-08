import "sort"

func isMajorityElement(nums []int, target int) bool {
	l := len(nums)
	b := sort.Search(l, func(i int) bool {
		return nums[i] >= target
	})
	e := sort.Search(l, func(i int) bool {
		return nums[i] > target
	})
	return (e-b)<<1 > l
}
