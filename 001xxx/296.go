import "sort"

func isPossibleDivide(nums []int, k int) bool {
	l := len(nums)
	if l%k != 0 {
		return false
	}
	m := make([]bool, l)
	sort.Ints(nums)
	s, next := 0, 1
	for s < l {
		n, k1, b := nums[s]+1, k-1, false
		m[s] = true
		i := next
		for ; i < l; i++ {
			if k1 == 0 {
				break
			} else if m[i] {
				continue
			} else if nums[i] == n {
				m[i], n, k1 = true, n+1, k1-1
				if !b {
					next, b = i+1, true
				}
			}
		}
		if k1 != 0 {
			return false
		}
		i, s = s, l
		for ; i < l; i++ {
			if !m[i] {
				s = i
				break
			}
		}
	}
	return true
}
