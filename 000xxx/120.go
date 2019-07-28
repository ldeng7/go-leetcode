func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	if 0 == l {
		return 0
	} else if 1 == l {
		return triangle[0][0]
	}
	t, s := make([]int, l), make([]int, l)
	t[0] = triangle[0][0]
	for i := 1; i < l; i++ {
		for j := 0; j <= i; j++ {
			s[j] = triangle[i][j]
			if j == 0 {
				s[j] += t[j]
			} else if j == i {
				s[j] += t[j-1]
			} else if t[j-1] < t[j] {
				s[j] += t[j-1]
			} else {
				s[j] += t[j]
			}
		}
		for j := 0; j <= i; j++ {
			t[j] = s[j]
		}
	}
	out := s[0]
	for i := 1; i < l; i++ {
		if s[i] < out {
			out = s[i]
		}
	}
	return out
}
