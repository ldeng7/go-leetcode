import "sort"

func threeSumSmaller(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	c := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[i]+nums[l]+nums[r] < target {
				c, l = c+r-l, l+1
			} else {
				r--
			}
		}
	}
	return c
}
