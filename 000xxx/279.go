import "math"

func numSquares(n int) int {
	t := make([]int, n+1)
	for i := 1; i <= n; i++ {
		t[i] = math.MaxInt64
	}
	for i := 0; i <= n; i++ {
		for j := 1; i+j*j <= n; j++ {
			if t[i]+1 < t[i+j*j] {
				t[i+j*j] = t[i] + 1
			}
		}
	}
	return t[len(t)-1]
}
