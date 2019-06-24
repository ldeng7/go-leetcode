import "sort"

func reversePairs(nums []int) int {
	var cal func(l, r int) int
	cal = func(l, r int) int {
		if l >= r {
			return 0
		}
		m := l + (r-l)>>1
		o := cal(l, m) + cal(m+1, r)
		for i, j := l, m+1; i <= m; i++ {
			for j <= r && nums[i] > nums[j]<<1 {
				j++
			}
			o += j - m - 1
		}
		sort.Ints(nums[l : r+1])
		return o
	}
	return cal(0, len(nums)-1)
}
