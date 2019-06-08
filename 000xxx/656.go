import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func cheapestJump(A []int, B int) []int {
	if -1 == A[len(A)-1] {
		return []int{}
	}
	n := len(A)
	t, p := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		t[i], p[i] = math.MaxInt64, -1
	}
	t[n-1] = A[n-1]

	for i := n - 2; i >= 0; i-- {
		if A[i] == -1 {
			continue
		}
		for j := i + 1; j <= min(i+B, n-1); j++ {
			if t[j] == math.MaxInt64 {
				continue
			}
			if A[i]+t[j] < t[i] {
				t[i], p[i] = A[i]+t[j], j
			}
		}
	}
	if t[0] == math.MaxInt64 {
		return []int{}
	}
	out := []int{}
	for i := 0; i != -1; i = p[i] {
		out = append(out, i+1)
	}
	return out
}
