import "sort"

var dirs = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type Node struct {
	nWalls  int
	walls   []int
	viruses []int
}

func containVirus(grid [][]int) int {
	o, m, n := 0, len(grid), len(grid[0])
	for {
		v := map[int]bool{}
		t := []Node{}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				k := (i << 8) | j
				if grid[i][j] != 1 || v[k] {
					continue
				}
				q, ws, vs := []int{k}, []int{}, []int{k}
				v[k] = true
				for 0 != len(q) {
					kt := q[0]
					q = q[1:]
					for _, d := range dirs {
						y, x := (kt>>8)+d[0], (kt&0xff)+d[1]
						k1 := (y << 8) | x
						if y < 0 || y >= m || x < 0 || x >= n || grid[y][x] == -1 || v[k1] {
							continue
						} else if grid[y][x] == 0 {
							ws = append(ws, k1)
						} else {
							vs, q, v[k1] = append(vs, k1), append(q, k1), true
						}
					}
				}
				m := map[int]bool{}
				for _, w := range ws {
					m[w] = true
				}
				t = append(t, Node{len(m), ws, vs})
			}
		}
		if 0 == len(t) {
			break
		}
		sort.Slice(t, func(i, j int) bool { return t[i].nWalls > t[j].nWalls })
		for _, k := range t[0].viruses {
			grid[k>>8][k&0xff] = -1
		}
		o += len(t[0].walls)
		for i := 1; i < len(t); i++ {
			for _, k := range t[i].walls {
				grid[k>>8][k&0xff] = 1
			}
		}
	}
	return o
}
