func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func totalFruit(tree []int) int {
	o, l, b, e := 0, len(tree), 0, 1
	s := map[int]bool{tree[0]: true}
	for i := 1; i < l; i++ {
		t := tree[i]
		if s[t] {
			e += 1
		} else if len(s) < 2 {
			s[t], e = true, e+1
		} else {
			o, e = max(e, o), i-b+1
			s = map[int]bool{t: true, tree[i-1]: true}
		}
		if t != tree[i-1] {
			b = i
		}
	}
	return max(e, o)
}
