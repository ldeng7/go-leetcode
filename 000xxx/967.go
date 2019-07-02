func numsSameConsecDiff(N int, K int) []int {
	s := map[int]bool{}
	for i := 1; i <= 9; i++ {
		s[i] = true
	}
	for i := 1; i <= N-1; i++ {
		s1 := map[int]bool{}
		for j, _ := range s {
			r := j % 10
			if r-K >= 0 {
				s1[j*10+r-K] = true
			}
			if r+K <= 9 {
				s1[j*10+r+K] = true
			}
		}
		s = s1
	}
	if N == 1 {
		s[0] = true
	}
	o, i := make([]int, len(s)), 0
	for j, _ := range s {
		o[i], i = j, i+1
	}
	return o
}
