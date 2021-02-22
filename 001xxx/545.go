func findKthBit(n int, k int) byte {
	if k == 1 {
		return '0'
	}
	m := 1 << (n - 1)
	if k == m {
		return '1'
	} else if k < m {
		return findKthBit(n-1, k)
	} else {
		k = m*2 - k
		return '0' + '1' - findKthBit(n-1, k)
	}
}
