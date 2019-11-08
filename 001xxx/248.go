func numberOfSubarrays(nums []int, k int) int {
	o, t := 0, make([]int, 1, len(nums))
	t[0] = 1
	for _, n := range nums {
		if n&1 != 0 {
			t = append(t, 0)
		}
		if len(t) > k {
			o += t[len(t)-1-k]
		}
		t[len(t)-1] += 1
	}
	return o
}
