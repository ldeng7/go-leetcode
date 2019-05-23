func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, num := range nums {
		if j, ok := m[num]; ok && i-j <= k {
			return true
		}
		m[num] = i
	}
	return false
}
