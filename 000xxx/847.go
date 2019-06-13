import "math"

func shortestPathLength(graph [][]int) int {
	l := uint16(len(graph))
	bse := uint16(1 << l)
	q := [][2]int{}
	t := make([][]int, l)
	for i := uint16(0); i < l; i++ {
		t[i] = make([]int, bse)
		for j := uint16(0); j < bse; j++ {
			t[i][j] = math.MaxInt64
		}
		t[i][1<<i] = 0
		q = append(q, [2]int{int(i), int(1 << i)})
	}

	for 0 != len(q) {
		q0 := q[0]
		i, bs := q0[0], q0[1]
		q = q[1:]
		for _, j := range graph[i] {
			bsj := bs | (1 << uint(j))
			if t[j][bsj] > t[i][bs]+1 {
				t[j][bsj] = t[i][bs] + 1
				q = append(q, [2]int{j, int(bsj)})
			}
		}
	}

	out, bs := math.MaxInt64, bse-1
	for i := uint16(0); i < l; i++ {
		if t[i][bs] < out {
			out = t[i][bs]
		}
	}
	return out
}
