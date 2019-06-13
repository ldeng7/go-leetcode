func constructArray(n int, k int) []int {
	out := []int{}
	i, j := 1, n
	for i <= j {
		var v int
		if k > 1 {
			if k&1 == 1 {
				v, i = i, i+1
			} else {
				v, j = j, j-1
			}
		} else {
			v, i = i, i+1
		}
		out, k = append(out, v), k-1
	}
	return out
}
