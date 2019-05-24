import "sort"

func triangleNumber(nums []int) int {
	sort.IntSlice(nums).Sort()
	c := 0
	for i := len(nums) - 1; i >= 2; i-- {
		j, k := 0, i-1
		for j < k {
			if nums[j]+nums[k] > nums[i] {
				c += k - j
				k--
			} else {
				j++
			}
		}
	}
	return c
}
