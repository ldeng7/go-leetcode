import "math"

func networkDelayTime(times [][]int, N int, K int) int {
	out := 0
	d := make([]int, N+1)
	for i := 0; i <= N; i++ {
		d[i] = math.MaxInt64
	}
	d[K] = 0
	for i := 1; i < N; i++ {
		for _, t := range times {
			if d[t[0]] != math.MaxInt64 && d[t[1]] > d[t[0]]+t[2] {
				d[t[1]] = d[t[0]] + t[2]
			}
		}
	}
	for i := 1; i <= N; i++ {
		if d[i] > out {
			out = d[i]
		}
	}
	if out == math.MaxInt64 {
		return -1
	}
	return out
}
