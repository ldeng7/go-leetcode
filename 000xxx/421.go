func findMaximumXOR(nums []int) int {
	o, m := 0, 0
	for i := 31; i >= 0; i-- {
		m |= 1 << uint(i)
		s := map[int]bool{}
		for _, n := range nums {
			s[n&m] = true
		}
		j := o | (1 << uint(i))
		for k, _ := range s {
			if s[j^k] {
				o = j
				break
			}
		}
	}
	return o
}
