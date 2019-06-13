func countPalindromicSubsequences(S string) int {
	n := len(S)
	t := make([][]int, n)
	for i := 0; i < n; i++ {
		t[i] = make([]int, n)
		t[i][i] = 1
	}
	for k := 1; k < n; k++ {
		for i := 0; i < n-k; i++ {
			j := i + k
			if S[i] == S[j] {
				l, r := i+1, j-1
				for l <= r && S[l] != S[i] {
					l++
				}
				for l <= r && S[r] != S[i] {
					r--
				}
				if l > r {
					t[i][j] = t[i+1][j-1]<<1 + 2
				} else if l == r {
					t[i][j] = t[i+1][j-1]<<1 + 1
				} else {
					t[i][j] = t[i+1][j-1]<<1 - t[l+1][r-1]
				}
			} else {
				t[i][j] = t[i][j-1] + t[i+1][j] - t[i+1][j-1]
			}
			if t[i][j] >= 0 {
				t[i][j] %= 1000000007
			} else {
				t[i][j] += 1000000007
			}
		}
	}
	return t[0][n-1]
}
