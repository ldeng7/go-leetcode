func minJump(jump []int) int {
	l := len(jump)
	t := make([]int, l)
	t[l-1] = 1
	for i := l - 2; i >= 0; i-- {
		if k := jump[i] + i; k < l {
			t[i] = t[k] + 1
		} else {
			t[i] = 1
		}
		for j := i + 1; j < l && t[j] >= t[i]+1; j++ {
			t[j] = t[i] + 1
		}
	}
	return t[0]
}
