func minimumTotal(triangle [][]int) int {
	if 0 == len(triangle) {
		return 0
	} else if 1 == len(triangle) {
		return triangle[0][0]
	}
	t := make([]int, len(triangle))
	s := make([]int, len(triangle))
	t[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
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
	for i := 1; i < len(triangle); i++ {
		if s[i] < out {
			out = s[i]
		}
	}
	return out
}
