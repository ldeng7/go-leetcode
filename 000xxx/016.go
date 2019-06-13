import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var diff *int
	out := 0

	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			d := sum - target
			if d < 0 {
				d = -d
			}
			if (nil != diff && d < *diff) || nil == diff {
				diff = &d
				out = sum
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return out
}
