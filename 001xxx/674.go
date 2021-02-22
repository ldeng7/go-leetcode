func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func minMoves(nums []int, limit int) int {
	l := len(nums)
	t := make([]int, (limit+1)*2)
	for i := 0; i < l/2; i++ {
		t[2] += 2
		a, b := nums[i], nums[l-i-1]
		t[min(a, b)+1]--
		t[a+b]--
		t[a+b+1]++
		t[max(a, b)+limit+1]++
	}
	o := 100000
	for i, s := 2, 0; i < limit*2+1; i++ {
		s += t[i]
		o = min(o, s)
	}
	return o
}
