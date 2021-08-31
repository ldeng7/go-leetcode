import "sort"

func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	o, l := 0, len(nums)
	for i, j := 0, l-1; i < l; i++ {
		for n := nums[i]; j > i; j-- {
			if n+nums[j] <= target {
				break
			}
		}
		if j > i {
			o += j - i
		}
	}
	return o % 1000000007
}
