import "sort"

var dirs = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func cutOffTree(forest [][]int) int {
	m, n := len(forest), len(forest[0])
	ts := make([][3]int, 0, 256)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if forest[i][j] > 1 {
				ts = append(ts, [3]int{forest[i][j], i, j})
			}
		}
	}
	sort.Slice(ts, func(i, j int) bool { return ts[i][0] < ts[j][0] })

	cal := func(r0, c0, r, c int) int {
		if r0 == r && c0 == c {
			return 0
		}
		cnt := 0
		q := [][2]int{{r0, c0}}
		v := make([][]bool, m)
		for i := 0; i < m; i++ {
			v[i] = make([]bool, n)
		}
		for 0 != len(q) {
			cnt++
			for i := len(q); i > 0; i-- {
				y0, x0 := q[0][0], q[0][1]
				q = q[1:]
				for _, d := range dirs {
					y, x := y0+d[0], x0+d[1]
					if y >= 0 && y < m && x >= 0 && x < n && !v[y][x] && forest[y][x] != 0 {
						if y == r && x == c {
							return cnt
						}
						v[y][x], q = true, append(q, [2]int{y, x})
					}
				}
			}
		}
		return -1
	}

	o, r, c := 0, 0, 0
	for _, t := range ts {
		cnt := cal(r, c, t[1], t[2])
		if cnt == -1 {
			return -1
		}
		o, r, c = o+cnt, t[1], t[2]
	}
	return o
}
