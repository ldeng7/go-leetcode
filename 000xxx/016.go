import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	o, l, d := 0, len(nums), math.MaxInt64
	sort.Ints(nums)
	for i := 0; i < l-2; i++ {
		j, k := i+1, l-1
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			d1 := s - target
			if d1 < 0 {
				d1 = -d1
			}
			if d1 < d {
				d, o = d1, s
			}
			if s > target {
				k--
			} else {
				j++
			}
		}
	}
	return o
}
