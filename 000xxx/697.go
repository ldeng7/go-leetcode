func findShortestSubArray(nums []int) int {
	mc, mb := map[int]int{}, map[int]int{}
	d, out := 0, len(nums)
	for i, n := range nums {
		mc[n]++
		if _, ok := mb[n]; !ok {
			mb[n] = i
		}
		if mc[n] == d {
			l := i - mb[n] + 1
			if l < out {
				out = l
			}
		} else if mc[n] > d {
			out, d = i-mb[n]+1, mc[n]
		}
	}
	return out
}
