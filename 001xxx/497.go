func canArrange(arr []int, k int) bool {
	cs := make([]int, k)
	for _, a := range arr {
		j := a % k
		if j < 0 {
			j += k
		}
		cs[j]++
	}
	if cs[0]&1 != 0 {
		return false
	}
	for i := 1; i < k/2; i++ {
		if cs[i] != cs[k-i] {
			return false
		}
	}
	return true
}
