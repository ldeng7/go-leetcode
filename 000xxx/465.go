import "math"

func minTransfers(transactions [][]int) int {
	m := map[int]int{}
	for _, t := range transactions {
		m[t[0]] -= t[2]
		m[t[1]] += t[2]
	}
	cs := make([]int, len(m))
	n := 0
	for _, v := range m {
		if v != 0 {
			cs[n] = v
			n++
		}
	}

	var cal func(int, int) int
	cal = func(b, num int) int {
		out := math.MaxInt64
		for b < n && cs[b] == 0 {
			b++
		}
		for i := b + 1; i < n; i++ {
			if (cs[i] < 0 && cs[b] > 0) || (cs[i] > 0 && cs[b] < 0) {
				cs[i] += cs[b]
				t := cal(b+1, num+1)
				if t < out {
					out = t
				}
				cs[i] -= cs[b]
			}
		}
		if out == math.MaxInt64 {
			return num
		}
		return out
	}
	return cal(0, 0)
}
