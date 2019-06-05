func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func candy(ratings []int) int {
	l := len(ratings)
	t := make([]int, l)
	for i := 0; i < l; i++ {
		t[i] = 1
	}
	for i := 0; i < l-1; i++ {
		if ratings[i+1] > ratings[i] {
			t[i+1] = t[i] + 1
		}
	}
	for i := l - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			t[i-1] = max(t[i-1], t[i]+1)
		}
	}
	o := 0
	for _, e := range t {
		o += e
	}
	return o
}
