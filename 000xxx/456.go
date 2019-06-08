import "math"

func find132pattern(nums []int) bool {
	third := math.MinInt64
	st := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < third {
			return true
		}
		for 0 != len(st) && nums[i] > st[len(st)-1] {
			third = st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, nums[i])
	}
	return false
}
