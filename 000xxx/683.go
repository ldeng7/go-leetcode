import "math"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func kEmptySlots(bulbs []int, K int) int {
	out, l, r, n := math.MaxInt64, 0, K+1, len(bulbs)
	days := make([]int, n)
	for i := 0; i < n; i++ {
		days[bulbs[i]-1] = i + 1
	}
	for i := 0; r < n; i++ {
		if days[i] < days[l] || days[i] <= days[r] {
			if i == r {
				t := max(days[l], days[r])
				if t < out {
					out = t
				}
			}
			l, r = i, K+i+1
		}
	}
	if out == math.MaxInt64 {
		return -1
	}
	return out
}
