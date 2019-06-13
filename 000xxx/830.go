func largeGroupPositions(S string) [][]int {
	bp, c, i, o := byte(' '), 0, 0, [][]int{}
	for ; i < len(S); i++ {
		if S[i] == bp {
			c++
		} else {
			if c >= 3 {
				o = append(o, []int{i - c, i - 1})
			}
			bp, c = S[i], 1
		}
	}
	if c >= 3 {
		o = append(o, []int{i - c, i - 1})
	}
	return o
}
