func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	m := map[int]int{}
	getKey := func(n int) int {
		if n >= 0 {
			return n / (t + 1)
		}
		return (n+1)/(t+1) - 1
	}
	for i, n := range nums {
		key := getKey(n)
		if _, ok := m[key]; ok {
			return true
		} else if v, ok := m[key-1]; ok && v >= n-t {
			return true
		} else if v, ok = m[key+1]; ok && v <= n+t {
			return true
		}
		m[key] = n
		if len(m) == k+1 {
			delete(m, getKey(nums[i-k]))
		}
	}
	return false
}
