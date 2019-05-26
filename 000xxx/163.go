import (
	"fmt"
)

func findMissingRanges(nums []int, lower int, upper int) []string {
	var rs [][2]int
	if 0 != len(nums) {
		if nums[0] != lower {
			rs = [][2]int{{lower, nums[0] - 1}}
		}
		for i := 1; i < len(nums); i++ {
			if nums[i]-nums[i-1] > 1 {
				rs = append(rs, [2]int{nums[i-1] + 1, nums[i] - 1})
			}
		}
		if nums[len(nums)-1] != upper {
			rs = append(rs, [2]int{nums[len(nums)-1] + 1, upper})
		}
	} else {
		rs = [][2]int{{lower, upper}}
	}
	out := make([]string, len(rs))
	for i, r := range rs {
		if r[0] != r[1] {
			out[i] = fmt.Sprintf("%d->%d", r[0], r[1])
		} else {
			out[i] = fmt.Sprintf("%d", r[0])
		}
	}
	return out
}
