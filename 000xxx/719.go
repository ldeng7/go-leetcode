import "sort"

func smallestDistancePair(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	l, r := 0, nums[n-1]-nums[0]
	for l < r {
		m := l + (r-l)>>1
		c, b := 0, 0
		for i := 0; i < n; i++ {
			for b < n && nums[i]-nums[b] > m {
				b++
			}
			c += i - b
		}
		if c < k {
			l = m + 1
		} else {
			r = m
		}
	}
	return r
}
