import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	o, m := 0, len(nums)>>1
	b := nums[m]
	for i := 0; i < m; i++ {
		o += b - nums[i]
	}
	for i := m + 1; i < len(nums); i++ {
		o += nums[i] - b
	}
	return o
}
