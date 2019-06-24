import "sort"

func largestDivisibleSubset(nums []int) []int {
	l, m, mi := len(nums), 0, 0
	o, t, p := []int{}, make([]int, l), make([]int, l)
	sort.Ints(nums)
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			if nums[j]%nums[i] == 0 && t[i] < t[j]+1 {
				t[i], p[i] = t[j]+1, j
				if m < t[i] {
					m, mi = t[i], i
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		o, mi = append(o, nums[mi]), p[mi]
	}
	return o
}
