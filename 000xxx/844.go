func backspaceCompare(S string, T string) bool {
	i, j, c1, c2 := len(S)-1, len(T)-1, 0, 0
	for ; i >= 0 || j >= 0; i, j = i-1, j-1 {
		for i >= 0 && (S[i] == '#' || c1 > 0) {
			if S[i] == '#' {
				c1++
			} else {
				c1--
			}
			i--
		}
		for j >= 0 && (T[j] == '#' || c2 > 0) {
			if T[j] == '#' {
				c2++
			} else {
				c2--
			}
			j--
		}
		if i < 0 || j < 0 {
			return i == j
		} else if S[i] != T[j] {
			return false
		}
	}
	return i == j
}
