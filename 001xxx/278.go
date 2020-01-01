import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func palindromePartition(s string, k int) int {
	l := len(s)
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([]int, l)
	}
	for i := 2; i <= l; i++ {
		for j := 0; j <= l-i; j++ {
			m := j + i - 1
			v := t[j+1][m-1]
			if s[j] != s[m] {
				v++
			}
			t[j][m] = v
		}
	}

	t1 := make([][]int, l+1)
	for i := 0; i <= l; i++ {
		ar := make([]int, k+1)
		for j := 0; j <= k; j++ {
			ar[j] = math.MaxInt64
		}
		t1[i] = ar
	}
	t1[0][0] = 0
	for i := 1; i <= l; i++ {
		t1[i][1] = t[0][i-1]
		je := min(k, i)
		for j := 2; j <= je; j++ {
			for m := j - 1; m < i; m++ {
				t1[i][j] = min(t1[i][j], t1[m][j-1]+t[m][i-1])
			}
		}
	}
	return t1[l][k]
}
