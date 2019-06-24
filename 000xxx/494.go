func findTargetSumWays(nums []int, S int) int {
	l := len(nums)
	t := make([]map[int]int, l+1)
	for i := 0; i <= l; i++ {
		t[i] = map[int]int{}
	}
	t[0][0] = 1
	for i := 0; i < l; i++ {
		n := nums[i]
		for s, c := range t[i] {
			t[i+1][s+n] += c
			t[i+1][s-n] += c
		}
	}
	return t[l][S]
}
