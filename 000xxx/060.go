func getPermutation(n int, k int) string {
	ds := make([]byte, 9)
	var i byte
	for ; i < 9; i++ {
		ds[i] = '1' + i
	}
	bs := make([]byte, n)
	t := make([]int, n)
	t[0] = 1
	for i := 1; i < n; i++ {
		t[i] = t[i-1] * i
	}
	k--
	for i := n; i >= 1; i-- {
		j := k / t[i-1]
		k %= t[i-1]
		bs[n-i] = ds[j]
		ds = append(ds[:j], ds[j+1:]...)
	}
	return string(bs)
}
