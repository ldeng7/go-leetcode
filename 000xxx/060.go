func getPermutation(n int, k int) string {
	digits := make([]byte, 9)
	var i byte
	for ; i < 9; i++ {
		digits[i] = '1' + i
	}
	bs := make([]byte, n)
	f := make([]int, n)
	f[0] = 1
	for i := 1; i < n; i++ {
		f[i] = f[i-1] * i
	}
	k--
	for i := n; i >= 1; i-- {
		j := k / f[i-1]
		k %= f[i-1]
		bs[n-i] = digits[j]
		digits = append(digits[:j], digits[j+1:]...)
	}
	return string(bs)
}
