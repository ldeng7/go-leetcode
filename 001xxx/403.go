import "sort"

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	o, s := make([]int, 0, len(nums)), 0
	for _, n := range nums {
		s += n
	}
	m := s >> 1
	s = 0
	for i := len(nums) - 1; i >= 0; i-- {
		n := nums[i]
		o = append(o, n)
		s += n
		if s > m {
			break
		}
	}
	return o
}
