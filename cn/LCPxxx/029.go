func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func orchestraLayout(n int, x int, y int) int {
	if x <= y {
		k := min(x, n-1-y)
		return (k*(n-k)*4+x+y-k*2)%9 + 1
	}
	k := min(y, n-1-x) + 1
	return (k*(n-k)*4-(x+y-(k-1)*2))%9 + 1
}
