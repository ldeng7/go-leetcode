import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	s := 0
	for _, n := range nums {
		s += n
	}
	if s%k != 0 {
		return false
	}
	t := s / k
	sort.Ints(nums)
	ar := make([]int, k)

	var cal func(int) bool
	cal = func(idx int) bool {
		if idx == -1 {
			for _, n := range ar {
				if n != t {
					return false
				}
			}
			return true
		}
		m := nums[idx]
		for i, n := range ar {
			if n+m > t {
				continue
			}
			ar[i] += m
			if cal(idx - 1) {
				return true
			}
			ar[i] -= m
		}
		return false
	}
	return cal(len(nums) - 1)
}
