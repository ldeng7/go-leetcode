import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, b := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		na, _ := strconv.Atoi(a + b)
		nb, _ := strconv.Atoi(b + a)
		return na < nb
	})
	if 0 == nums[len(nums)-1] {
		return "0"
	}
	s := ""
	for i := len(nums) - 1; i >= 0; i-- {
		s += strconv.Itoa(nums[i])
	}
	return s
}
