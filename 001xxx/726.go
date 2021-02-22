func tupleSameProduct(nums []int) int {
	l := len(nums)
	m := make(map[int]int, l*l)
	for i, n := range nums {
		for j := i + 1; j < l; j++ {
			m[n*nums[j]]++
		}
	}
	o := 0
	for _, c := range m {
		if c > 1 {
			o += ((c * (c - 1)) >> 1) << 3
		}
	}
	return o
}
