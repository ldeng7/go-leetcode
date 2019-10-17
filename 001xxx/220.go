const M = 1000000007

func countVowelPermutation(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1
	for j := 2; j <= n; j++ {
		a, e, i, o, u = (e+i+u)%M, (a+i)%M, (e+o)%M, i, (o+i)%M
	}
	return (a + e + i + o + u) % M
}
