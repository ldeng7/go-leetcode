import "sort"

func minMalwareSpread(graph [][]int, initial []int) int {
	sort.Ints(initial)
	o, l, m := initial[0], len(graph), 0
	v, s := make([]bool, l), map[int]bool{}
	for _, i := range initial {
		s[i] = true
	}

	var cal func(int, int) int
	cal = func(i, n int) int {
		for j := 0; j < l; j++ {
			if v[j] || graph[i][j] != 1 || j == i || n <= 0 {
				continue
			}
			v[j] = true
			if s[j] {
				n = -1
				break
			}
			n = cal(j, n+1)
		}
		return n
	}

	for _, i := range initial {
		if v[i] {
			continue
		}
		v[i] = true
		n := cal(i, 1)
		if n < 0 {
			break
		}
		if n > m {
			o, m = i, n
		}
	}
	return o
}
