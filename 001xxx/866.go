var t [1001][1001]int

func rearrangeSticks(n int, k int) int {
	if k == n {
		return 1
	} else if n == 0 || k == 0 {
		return 0
	} else if 0 != t[n][k] {
		return t[n][k]
	}
	o := ((n-1)*rearrangeSticks(n-1, k) + rearrangeSticks(n-1, k-1)) % 1000000007
	t[n][k] = o
	return o
}
