import "fmt"

func summaryRanges(nums []int) []string {
	out := []string{}
	i, l := 0, len(nums)
	for i < l {
		j := 1
		for i+j < l && nums[i+j]-nums[i] == j {
			j++
		}
		var s string
		if j > 1 {
			s = fmt.Sprintf("%d->%d", nums[i], nums[i+j-1])
		} else {
			s = fmt.Sprintf("%d", nums[i])
		}
		out = append(out, s)
		i += j
	}
	return out
}
