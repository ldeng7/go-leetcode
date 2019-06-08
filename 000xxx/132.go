import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minCut(s string) int {
	l := len(s)
	if 0 == l {
		return 0
	}
	t := make([]int, l+1)
	t[0] = -1
	for i := 1; i <= l; i++ {
		t[i] = math.MaxInt64
	}
	for i := 0; i < l; i++ {
		for j := 0; i-j >= 0 && i+j < l && s[i-j] == s[i+j]; j++ {
			t[i+j+1] = min(t[i+j+1], t[i-j]+1)
		}
		for j := 0; i-j >= 0 && i+j+1 < l && s[i-j] == s[i+j+1]; j++ {
			t[i+j+2] = min(t[i+j+2], t[i-j]+1)
		}
	}
	return t[l]
}
