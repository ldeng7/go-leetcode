import "sort"

func subsetsWithDup(nums []int) [][]int {
	out := [][]int{{}}
	if 0 == len(nums) {
		return out
	}
	sort.Ints(nums)
	sz, last := 1, nums[0]
	for _, num := range nums {
		if num != last {
			last = num
			sz = len(out)
		}
		sz1 := len(out)
		for i := sz1 - sz; i < sz1; i++ {
			ar1 := make([]int, len(out[i])+1)
			copy(ar1, out[i])
			ar1[len(out[i])] = num
			out = append(out, ar1)
		}
	}
	return out
}
