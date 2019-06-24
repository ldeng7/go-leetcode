func circularArrayLoop(nums []int) bool {
	l := len(nums)
	bs := make([]bool, l)
	for i := 0; i < l; i++ {
		if bs[i] {
			continue
		}
		bs[i] = true
		m, j := map[int]int{}, i
		for {
			k := (j + nums[j]) % l
			if k < 0 {
				k += l
			}
			if k == j || nums[k]*nums[j] < 0 {
				break
			}
			if _, ok := m[k]; ok {
				return true
			}
			m[j], j, bs[k] = k, k, true
		}
	}
	return false
}
