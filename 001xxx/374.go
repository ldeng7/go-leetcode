func generateTheString(n int) string {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = 'a'
	}
	if n&1 == 0 {
		bs[n-1] = 'b'
	}
	return string(bs)
}
