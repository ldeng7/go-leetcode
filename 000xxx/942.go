func diStringMatch(S string) []int {
	l := len(S)
	m := map[int]bool{}
	out := make([]int, l+1)

	cnt := 0
	for _, c := range S {
		if c == 'I' {
			cnt++
		}
	}
	out[l] = cnt
	m[cnt] = true

	for i := 0; i < l; i++ {
		if S[l-i-1] == 'I' {
			for j := out[l-i] - 1; j >= 0; j-- {
				if _, ok := m[j]; !ok {
					out[l-i-1], m[j] = j, true
					break
				}
			}
		} else {
			for j := out[l-i] + 1; j <= l; j++ {
				if _, ok := m[j]; !ok {
					out[l-i-1], m[j] = j, true
					break
				}
			}
		}
	}
	return out
}
