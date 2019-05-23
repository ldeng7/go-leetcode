func findJudge(N int, trust [][]int) int {
	er := make([]bool, N)
	ee := make([]int, N)
	for _, t := range trust {
		a, b := t[0]-1, t[1]-1
		er[a] = true
		ee[b] += 1
	}
	for i, e := range er {
		if !e && ee[i] == N-1 {
			return i + 1
		}
	}
	return -1
}
