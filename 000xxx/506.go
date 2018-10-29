import "sort"
import "strconv"

func findRelativeRanks(nums []int) []string {
	out := make([]string, len(nums))
	cnt := 1
	m := map[int]int{}
	for i, n := range nums {
		m[n] = i
	}
	sort.IntSlice(nums).Sort()
	for i := len(nums) - 1; i >= 0; i-- {
		j := m[nums[i]]
		if cnt == 1 {
			out[j] = "Gold Medal"
		} else if cnt == 2 {
			out[j] = "Silver Medal"
		} else if cnt == 3 {
			out[j] = "Bronze Medal"
		} else {
			out[j] = strconv.Itoa(cnt)
		}
		cnt++
	}
	return out
}
